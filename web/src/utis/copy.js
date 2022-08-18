export default function copyCommand(text) {
  if(!text){
    return
  }
  return new Promise((resolve, reject) => {
    try{
      var input = document.createElement("input"); // 直接构建input
      input.style.position = 'fixed'
      input.style.left = -window.screen.width *  2 + 'px';
      input.style.opacity = 0
      input.value = text; // 设置内容
      document.body.appendChild(input); // 添加临时实例
      input.select(); // 选择实例内容
      document.execCommand("Copy"); // 执行复制
      document.body.removeChild(input); // 删除临时实例
      resolve()
    } catch {
      reject()
    }
  })
}
