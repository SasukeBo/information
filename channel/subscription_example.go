package channel

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Post struct {
	ID    int `json:"id"`
	Likes int `json:"count"`
}

type ConnectionACKMessage struct {
	OperationID string `json:"id,omitempty"`
	Type        string `json:"type"`
	Payload     struct {
		Query string `json:"query"`
	} `json:"payload,omitempty"`
}

var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"likes": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

type Subscriber struct {
	ID            int
	Conn          *websocket.Conn
	RequestString string
	OperationID   string
}

func main() {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{"graphql-ws"},
	}
	var posts = []*Post{
		&Post{ID: 1, Likes: 1},
		&Post{ID: 2, Likes: 2},
	}
	var subscribers sync.Map
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"posts": &graphql.Field{
					Type: graphql.NewList(PostType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return posts, nil
					},
				},
			},
		}),
		Subscription: graphql.NewObject(graphql.ObjectConfig{
			Name: "Subscription",
			Fields: graphql.Fields{
				"postLikesSubscribe": &graphql.Field{
					Type: graphql.NewList(PostType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return posts, nil
					},
				},
			},
		}),
	})
	if err != nil {
		log.Fatal(err)
	}
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})
	http.Handle("/graphql", h)
	http.HandleFunc("/subscriptions", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("failed to do websocket upgrade: %v", err)
			return
		}
		connectionACK, err := json.Marshal(map[string]string{
			"type": "connection_ack",
		})
		if err != nil {
			log.Printf("failed to marshal ws connection ack: %v", err)
		}
		if err := conn.WriteMessage(websocket.TextMessage, connectionACK); err != nil {
			log.Printf("failed to write to ws connection: %v", err)
			return
		}
		go func() {
			for {
				_, p, err := conn.ReadMessage()
				if websocket.IsCloseError(err, websocket.CloseGoingAway) {
					return
				}
				if err != nil {
					log.Println("failed to read websocket message: %v", err)
					return
				}
				var msg ConnectionACKMessage
				if err := json.Unmarshal(p, &msg); err != nil {
					log.Printf("failed to unmarshal: %v", err)
					return
				}
				if msg.Type == "start" {
					length := 0
					subscribers.Range(func(key, value interface{}) bool {
						length++
						return true
					})
					var subscriber = Subscriber{
						ID:            length + 1,
						Conn:          conn,
						RequestString: msg.Payload.Query,
						OperationID:   msg.OperationID,
					}
					subscribers.Store(subscriber.ID, &subscriber)
				}
			}
		}()
	})
	go func() {
		for {
			time.Sleep(1 * time.Second)
			for _, post := range posts {
				post.Likes = post.Likes + 1
			}
			subscribers.Range(func(key, value interface{}) bool {
				subscriber, ok := value.(*Subscriber)
				if !ok {
					return true
				}
				payload := graphql.Do(graphql.Params{
					Schema:        schema,
					RequestString: subscriber.RequestString,
				})
				message, err := json.Marshal(map[string]interface{}{
					"type":    "data",
					"id":      subscriber.OperationID,
					"payload": payload,
				})
				if err != nil {
					log.Printf("failed to marshal message: %v", err)
					return true
				}
				if err := subscriber.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
					if err == websocket.ErrCloseSent {
						subscribers.Delete(key)
						return true
					}
					log.Printf("failed to write to ws connection: %v", err)
					return true
				}
				return true
			})
		}
	}()
	log.Printf("server running on port :8080")
	http.ListenAndServe(":8080", nil)
}
