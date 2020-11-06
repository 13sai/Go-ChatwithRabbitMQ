<template>
  <div class="message">
    <van-nav-bar
      title="标题"
      class="nav"
    />
    <ul class="message-ul">
      <li v-for="item in list">
          <!-- <p class="time">
              <span>{{ item.date | time }}</span>
          </p> -->
          <div class="main" :class="{ self: item.self }">
              <img class="avatar" src="./../../assets/logo.png" />
              <div class="text">{{ item.message }}</div>
          </div>
      </li>
    </ul>
</div>
</template>

<script>
export default {
  data() {
    return {
      username: this.$store.getters.name,
      password: '',
      finished: false,
      list: [
        
      ]
    };
  },
  created() {
      this.initWebSocket();
    },
    destroyed() {
      this.websock.close() //离开路由之后断开websocket连接
    },
    methods: {
      initWebSocket(){ //初始化weosocket
        const wsuri = "ws://127.0.0.1:9003/ws";
        this.websock = new WebSocket(wsuri);
        this.websock.onmessage = this.websocketonmessage;
        this.websock.onopen = this.websocketonopen;
        this.websock.onerror = this.websocketonerror;
        this.websock.onclose = this.websocketclose;
      },
      websocketonopen(){ //连接建立之后执行send方法发送数据
        // let actions = {"test":"12345"};
        // this.websocketsend(JSON.stringify(actions));
      },
      websocketonerror(){//连接建立失败重连
        this.initWebSocket();
      },
      websocketonmessage(e){ //数据接收
        console.log(e.data)
        const redata = JSON.parse(e.data);
        this.list.push(redata)
      },
      websocketsend(Data){//数据发送
        this.websock.send(Data);
      },
      websocketclose(e){  //关闭
        console.log('断开连接',e);
      },
    },
};
</script>

<style>
.nav {
  background: #b2e281;
  color: #fff;
}
.chat-container{
    position: relative;
    width: 100%;
    height: 100%;
    background: #999;
}
.message {
  overflow-y: scroll;
}
.message-ul {
  margin-top: 2rem;
}
.message li {
  margin: 1rem;
}
.message li .main {
  text-align: left;
}
.message .time {
  margin: 7rem 0;
  text-align: center;
}
.message .time > span {
  display: inline-block;
  padding: 0 18rem;
  font-size: 12rem;
  color: #fff;
  border-radius: 2rem;
  background-color: #dcdcdc;
}
.message .avatar {
  float: left;
  margin: 0 1.5rem 0 0;
  border-radius: 3rem;
  width: 3rem;
}
.message .text {
  display: inline-block;
  position: relative;
  padding: 0 1rem;
  min-height: 4rem;
  font-size: 3rem;
  text-align: left;
  word-break: break-all;
  background-color: #cdcdcd;
  border-radius: 1rem;
}
.message .text:before {
  content: " ";
  position: absolute;
  top: 1rem;
  right: 100%;
  border: 1rem solid transparent;
  border-right-color: #cdcdcd;
}
.message .self {
  text-align: right;
}
.message .self .avatar {
  float: right;
  margin: 0 0 0 10rem;
}
.message .self .text {
  background-color: #b2e281;
}
.message .self .text:before {
  right: inherit;
  left: 100%;
  border-right-color: transparent;
  border-left-color: #b2e281;
}
</style>