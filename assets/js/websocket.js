class Channel {
  constructor(channel, opts, socket) {
    this.channel = channel
    this.opts = opts
    this.socket = socket
    this._onData = null
    this.topics = []
  }

  Join(topic) {
    if (!this.topics.includes(topic)) this.topics.push(topic)
    var message = {
      channel: this.channel,
      event: 'join',
      topic
    }

    this.socket.send(JSON.stringify(message))
  }

  set onData(func) {
    this._onData = func
  }

  Leave(topic) {
    var index = this.topics.indexOf(topic)
    if (index > -1) this.topics.splice(index, 1)
    var message = {
      channel: this.channel,
      event: 'leave',
      topic
    }

    this.socket.send(JSON.stringify(message))
  }

  Send(topic, payload) {
    var message = {
      channel: this.channel,
      event: 'data',
      payload,
      topic
    }

    this.socket.send(JSON.stringify(message))
  }
}

class Socket {
  constructor(endPoint) {
    this.socketUrl = `ws://${document.location.hostname}${endPoint}`
    this.socket = null
    this.HartInterval = null
    this.channels = {}
  }

  connect() {
    console.log('WebSocket connecting...')
    this.socket = new WebSocket(this.socketUrl)
    this.socket.onopen = () => {
      console.log('WebSocket is open now.')
      this.rejoin()
      this.HartCheck()
    }
    this.socket.onclose = () => {
      console.log('WebSocket is closed now. try to reconnect')
      var reconnInterval = setInterval(() => {
        if (this.socket.readyState == this.socket.OPEN) {
          clearInterval(reconnInterval)
          return
        }
        if (this.socket.readyState == this.socket.CLOSED) this.connect()
      }, 5000)
    }
    this.socket.onmessage = ({ data: dataStr }) => {
      var data = JSON.parse(dataStr)
      console.log('onmessage channel ', data.channel)
      var channel = this.channels[data.channel]
      if (channel && channel._onData) {
        channel._onData(data)
      }
    }
  }

  rejoin() {
    for (var chanName in this.channels) {
      var channel = this.channels[chanName]
      channel.socket = this.socket
      channel.topics.forEach(topic => {
        channel.Join(topic)
      })
    }
  }

  close() {
    this.socket.send(`{"channel": "close"}`)
  }

  channel(channel, opts) {
    if (this.channels[channel]) return this.channels[channel]
    this.channels[channel] = new Channel(channel, opts, this.socket)
    return this.channels[channel]
  }

  HartCheck() {
    clearInterval(this.HartInterval)
    this.HartInterval = setInterval(() => {
      this.socket.send(`{"channel": "heartbeat", "payload": {"message": "ping"}}`)
    }, 9000)
  }
}

export default Socket
