<template>
  <div class="message">
    <van-nav-bar
      :title="clientName"
      class="nav"
      left-text="返回"
      left-arrow
      @click-left="onClickLeft"
    />
    <ul class="message-ul">
        <li v-for="(item, i) in list" :key="i">
            <!-- <p class="time">
                <span>{{ item.date | time }}</span>
            </p> -->
            <div class="main2" v-if="item.mine">
                <div class="text">{{ item.message }}</div>
                <img class="avatar" :src="avatar"/>
            </div>
            <div class="main" v-else>
                <div class="text">{{ item.message }}</div>
                <img class="avatar" :src="clientAvatar"/>
            </div>
        </li>
    </ul>
    <div class="inputBox">
        <van-cell-group>
            <van-field
                v-model="value"
                rows="2"
                autosize
                type="input"
                placeholder="请输入留言"
            >
                <template #button>
                    <van-button size="small" type="primary" @click="send">发送</van-button>
                </template>
            </van-field>
        </van-cell-group>
    </div>
</div>
</template>

<script>
import { getToken } from "@/utils/auth";

export default {
  data() {
    return {
      avatar: this.$store.getters.avatar,
      finished: false,
      value: '',
      clientId: this.$route.params.id,
      clientAvatar: this.$route.query.avatar,
      clientName: this.$route.query.name,
      mineId: '',
      list: [
      ]
    };
  },
  created() {
    this.initWebSocket();
    console.log(this.$route.params.id)
  },
  destroyed() {
    this.websock.close() //离开路由之后断开websocket连接
  },
  methods: {
    initWebSocket(){ //初始化weosocket
      const wsuri = "ws://127.0.0.1:9003/ws?token="+getToken();
      this.websock = new WebSocket(wsuri);
      this.websock.onmessage = this.websocketonmessage;
      this.websock.onopen = this.websocketonopen;
      this.websock.onerror = this.websocketonerror;
      this.websock.onclose = this.websocketclose;
    },
    send() {
      this.websocketsend();
    },
    websocketonopen(){ //连接建立之后执行send方法发送数据
    },
    websocketonerror(){//连接建立失败重连
      this.initWebSocket();
    },
    websocketonmessage(e){ //数据接收
      const redata = JSON.parse(e.data);
      this.list.push(redata)
    },
    websocketsend(){//数据发送
      if (this.value.length < 1) {
          this.$toast("请输入内容");
          return false;
      }
      let actions = {"clientId":this.clientId, "message": this.value};
      this.websock.send(JSON.stringify(actions));
      this.value = "";
    },
    websocketclose(e){  //关闭
      console.log('断开连接');
    },
    onClickLeft() {
      this.$router.push({path: "/chat/list"})
    },
  },
};
</script>

<style>
.message .nav {
  background: #07c160;
}
.message .nav .van-nav-bar__title, .van-nav-bar__text,  .van-nav-bar__arrow {
  color: #fff!important;
}
.chat-container{
    position: relative;
    width: 100%;
    height: 100%;
}
.message {
    overflow-y: scroll;
}
.message-ul {
    margin-top: 10px;
}
.message li {
  margin: 5px;
  font-size: 20px;
}
.message .main {
    text-align: left;
}
.message .main .avatar {
    width: 30px;
    height: 30px;
    float: left;
}
.message .main2 {
    text-align: right;
}
.message .main2 .avatar {
    width: 30px;
    height: 30px;
    float: right;
}
.message .main2 .text {
    position: relative;
    margin-right: 10px;
    background: #fff;
    border-radius: 4px;
    padding: 4px;
    display: inline-block;
    word-break: break-all; 
}
.message .main2 .text::before {
    content: " ";
    position: absolute;
    top: 5px;
    right: -10px;
    border: 6px solid transparent;
    border-left-color: #fff;
}
.message .main .text {
    position: relative;
    margin-left: 10px;
    background: #fff;
    border-radius: 4px;
    padding: 4px;
    display: inline-block;
    word-break: break-all; 
}
.message .main .text::before {
    content: " ";
    position: absolute;
    top: 5px;
    right: 100%;
    border: 6px solid transparent;
    border-right-color: #fff;
}

.message .inputBox {
    position: fixed;
    bottom: 0;
    width: 100%;
    max-width: 375px;
}
</style>