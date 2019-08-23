class Channel {
  constructor(channel, opts, socket) {
    this.channel = channel // channel 名称
    this.opts = opts // 其它选项
    this.socket = socket // 实际的socket连接
    this._onData = null // channel 接收消息处理的回调函数
    this.topics = [] // 当前channel加入的topic，用于连接恢复时重新加入topic
  }

  // channel 加入话题
  Join(topic) {
    if (!this.topics.includes(topic)) this.topics.push(topic)
    var message = {
      channel: this.channel,
      event: 'join',
      topic
    }

    this.socket.send(JSON.stringify(message))
  }

  // channel 消息处理回调函数
  set onData(func) {
    this._onData = func
  }

  // channel 离开话题
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

  // channel 发送数据
  Send(topic, payload) {
    var message = {
      channel: this.channel,
      event: 'data',
      payload,
      topic
    }

    this.socket.send(JSON.stringify(message))
  }

  //
  destroy() {

  }
}

class Socket {
  constructor(endPoint) {
    this.socketUrl = `ws://${document.location.hostname}${endPoint}` // websocket 服务器链接
    this.socket = null // 实际的socket
    this.HartInterval = null // 心跳间隔
    this.channels = {} // socket 中的 channel
    this.opts = null
  }

  // 发起 websocket 连接
  connect(opts) {
    this.opts = opts
    console.log('WebSocket connecting...')
    this.socket = new WebSocket(this.socketUrl)
    this.socket.onopen = () => {
      console.log('WebSocket is open now.')
      var message = {
        channel: 'system',
        ...opts
      }
      this.socket.send(JSON.stringify(message))
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
        if (this.socket.readyState == this.socket.CLOSED) this.connect(this.opts)
      }, 5000)
    }
    this.socket.onmessage = ({ data: dataStr }) => {
      var data = JSON.parse(dataStr)
      if (data.channel === 'system' && data.topic === 'error') {
        console.error(data.payload)
        this.close()
      }

      var channel = this.channels[data.channel]
      if (channel && channel._onData) {
        channel._onData(data)
      }
    }
  }

  // 重新加入话题
  rejoin() {
    for (var chanName in this.channels) {
      var channel = this.channels[chanName]
      channel.socket = this.socket
      channel.topics.forEach(topic => {
        channel.Join(topic)
      })
    }
  }

  // 关闭 socket
  close() {
    clearInterval(this.HartInterval) // 关闭心跳
    this.socket.onclose = undefined // 删除断连回调
    this.socket.close() // 关闭 socket
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
