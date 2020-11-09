<template>
  <div class="chat-list">
    <van-nav-bar
      title="列表"
      class="nav"
    />
    <van-list
      v-model="loading"
      :finished="finished"
      finished-text=""
      @load="getList"
    >
      <div class="chat-list-div" v-for="(item, i) in list" :key="i" @click="chat(item)">
        <van-image :src="item.avatar"/>
        <span>{{item.nickname}}</span>
      </div>
    </van-list>
</div>
</template>

<script>
import { getToken } from "@/utils/auth";
import { getUserList } from "@/api/user";

export default {
  data() {
    return {
      username: this.$store.getters.name,
      finished: false,
      loading: false,
      value: '',
      clientId: '',
      query: {
        limit: 10,
        page: 1
      },
      mineId: '',
      list: [
        
      ]
    };
  },
  created() {
    // this.getList();
    this.initWebSocket();
  },
  destroyed() {
    this.websock.close() //离开路由之后断开websocket连接
  },
  methods: {
    getList() {
      getUserList(this.query).then(res => {
        var data = res.data;
        console.warn(data)
        this.list = data.data;
        this.query.page = data.page;
        if (data.page >= data.total_page) {
          this.finished = true;
        }
        this.loading = false;
      })
    },
    chat(row) {
      this.$router.push({path: "/chat/detail/"+row.id, query: {avatar: row.avatar, name: row.nickname}})
    },
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
      // this.list.push(redata)
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
.chat-list .nav {
  background: #07c160;
}
.chat-list .nav .van-nav-bar__title {
  color: #fff!important;
}
.chat-list-div {
  width: 100%;
  background: #fff;
  border-bottom: 0.1rem solid #ddd;
  height: 60px;
  line-height: 60px; 
  font-size: 20px;
  text-align: left;
}
.chat-list-div span {
  margin-left: 10px;
  font-weight: bold;
}
.chat-list-div .van-image {
  float: left;
  width: 40px;
  height: 40px;
  margin: 10px 0 0 10px;
}
.chat-list-div .van-image img {
  border-radius: 5px;
}
</style>