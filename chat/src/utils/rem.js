function setRem () {
  let htmlWidth = document.documentElement.clientWidth ||document.body.clientWidth;  //检测html的屏幕宽度和body的屏幕宽度
  document.documentElement.style.fontSize= htmlWidth/7.5 + 'px'; //设置font-size在750屏幕下的尺寸为100px,这样转换的rem可以一目了然之前是多少px的。开发屏幕尺寸自己选，3.2的320屏幕宽也可以。
}

setRem();
window.onresize = function () {   //浏览器尺寸变化进行触发window.onresize函数，然后触发函数setRem()
  setRem()
}