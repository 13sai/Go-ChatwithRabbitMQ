<template>
  <div class="message">
    <van-nav-bar
      title="开放聊天室"
      class="nav"
    />
    <ul class="message-ul">
        <li v-for="item in list">
            <!-- <p class="time">
                <span>{{ item.date | time }}</span>
            </p> -->
            <div class="main2" v-if="item.mine">
                <div class="text">{{ item.message }}</div>
                <img class="avatar" src="https://ss0.bdstatic.com/70cFuHSh_Q1YnxGkpoWK1HF6hhy/it/u=2422475850,1168377142&fm=15&gp=0.jpg" />
            </div>
            <div class="main" v-else>
                <div class="text">{{ item.message }}</div>
                <img class="avatar" src="https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=1128428061,2613189316&fm=15&gp=0.jpg" />
            </div>
        </li>
    </ul>
    <div class="inputBox">
        <van-cell-group>
            <van-field
                v-model="value"
                rows="2"
                autosize
                type="textarea"
                maxlength="50"
                placeholder="请输入留言"
                show-word-limit
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
      username: this.$store.getters.name,
      finished: false,
      value: '',
      clientId: '',
      mineId: '',
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
        console.log('断开连接',e);
      },
    },
};
</script>

<style>
.message .nav {
  background: #07c160;
}
.message .nav .van-nav-bar__title {
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
    margin-top: 2rem;
}
.message li {
  margin: 1rem;
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
    line-height: 24px;
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
    line-height: 24px;
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
}
</style>