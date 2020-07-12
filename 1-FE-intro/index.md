# 前端再识 &amp; HTML5 API

## 前端再识 & HTML5 API


### 前端再识


#### 前端与新技术的碰撞


- AI



AI 不重 UI ，所以前端技术并不会给 AI 带来多大贡献，使用时也是由后端对数据进行处理，前端只关心数据结果。TensorFlow.js 虽然基于 JavaScript，但也只是因为 JavaScript 在做机器学习时有它的优势，JavaScript 无法代表前端技术。


- VR/AR



Web 端已经有相对成熟的解决方案来展示 VR/AR（ React360 、 AR.js 、 WebRTC 等），我们可以通过浏览器看到 VR/AR 的内容，因此这是一个可能的方向，但同时也要知道， VR/AR 的发展一直都重在设备而不在展示内容，没有完善的硬件植入， VR/AR 永远也达不到这项技术最理想的效果。


- IoT



IoT 带来的是交互设备的改变，一个又一个的新型终端会逐渐影响原来无线时代智能手机的统治地位，Web 的高度可移植性会让它很快在各种设备中生根发芽。有个说法是IoT是前端的新蓝海，确实不无道理，IoT时代，水平稍低的前端可以开发各种平台的应用，水平较高的前端可以通过移植浏览器内核，打造IoT设备的前端Runtime。


- 5G



5G 给人最直观的感受是网速的提升，也许我们今天考虑的很多性能优化在 5G 时代不再需要，云服务会越来越兴盛，serverless 逐渐成为主流，甚至渲染也会转移到云端，本来限于浏览器性能而发展缓慢的 WebGL 等技术也许会展现活力。


#### 前端岗位需要具备的责任感


前端工程师在开发中所处的位置大家基本已经全面了解，承上启下，承接上层需求，对接下层数据，需要懂产品、会设计、精开发，在一个项目从想法到落地的过程中，你要能


- 理解产品需求



1. 对于不合理的地方提出自己的意见，或砍去或修改
2. 评估项目难度，不能做的或是难度配不上价值的，主动放弃，能做的给出合适的项目周期
3. 对于产品没有想到的功能，勇于提出，给项目增加合理的需求



- 完美还原视觉稿



1. 成品图和视觉稿在颜色上不要有一位的偏差，如果觉得颜色有必要修改，和视觉商量
2. 在各种设备上的体验都要尽可能趋于一致，如果在某型号设备上偏差过大，考虑单独给出一版
3. 对于视觉要求的动画效果能做的全部做到，不能做的、难度大的、性能消耗大的、无意义的给予驳回
4. 主动考虑怎么优化用户的交互体验，尝试自己增加一些过渡效果



- 高效率开发



1. 如果开发者不止你一个，善于使用合作开发工具并分配好各自任务
2. 要尝试自己去指定接口格式，最不济也要和后端商量指定，不要给什么用什么，不好用或是不合理仍然去用
3. 学会利用 IDE 减少自己的非脑力劳动
4. 采用已有的类库、框架去开发项目，不要重复造轮子，想写轮子练手在业余时间
5. 不要盲目采用新技术，要看项目是否需要，想学习新技术在业余时间



**身为开发者，尤其是前端开发者，做出的东西一定要自己觉得不错，别人觉得好用**


#### 前端技能点


##### 基本技能


- HTML 和 CSS
- JavaScript 基础
- 基本的 http 协议
- 移动端适配
- Git 版本控制与合作开发模式
- 使用 node 搭建开发环境
- 使用包管理工具（ npm / yarn ）
- 使用 webpack 进行打包构建
- 使用至少一种框架进行快速开发



##### 进阶要求


- 基本的计算机网络知识
- 基本的操作系统概念
- 基本的 Linux 操作
- 对各种数据结构都有了解并能用至少一种语言实现
- 理解常见的架构模式与设计模式
- 掌握一种编程范式
- 掌握 ES6+ 语法
- 掌握 HTML5 和 CSS3 ，了解最新 W3C 标准
- 熟练使用至少一种 CSS 预处理器
- 使用 npm scripts （次选 gulp ）完成前端自动化
- 掌握 Lint 与代码格式化工具
- 能使用 webpack 独立完成 vue / react 脚手架的大部分功能
- 理解常用框架的基本实现原理
- 了解浏览器运行原理
- 了解至少一门后端语言并能用来做小型项目开发
- 手动搭建一个简单的 nginx 服务器
- 学会利用缓存优化网站性能
- 了解常见的网络攻击手段并能够在前端过滤
- 了解基本的 SEO
- 掌握至少一种跨端解决方案
- 熟练使用调试工具



##### 持续学习


- 开始使用 TypeScript
- 阅读框架源码
- WebGL 图形学、Canvas可视化
- BFF 、serverless、nodejs中间层构建
- 关注 WebAssembly
- 小程序
- 学习使用各种运维工具
- 学习 PWA（渐进式 Web 应用）
- 学习一种自动化测试框架
- 积极参与开源项目
- 了解互联网
- 保持身体健康



### HTML5


一方面它定义一个新版本语言，具有新的元素、属性和行为；另一方面它有更大的技术集，允许构建更多样化和更强大的网站和应用程序。下面列出 HTML5 新增的特性：


- 语义：能够让你更恰当地描述你的内容是什么，使得页面更易读，结构更清晰。

```html
<header> 、<footer>、<hgroup>、<nav>、<article>、<aside>、<address>
```

- 连通性：能够让你和服务器之间通过创新的新技术方法进行通信



> WebSockets ...



- 离线 & 存储：能够让网页在客户端本地存储数据以及更高效地离线运行
- 多媒体：使 video 和 audio 成为了在所有 Web 中的一等公民
- 2D/3D 绘图 & 效果：提供了一个更加分化范围的呈现选择



> Canvas、WegGL、SVG...



- 性能 & 集成：提供了非常显著的性能优化和更有效的计算机硬件使用



> Web Workers、History API、拖拽...



- 设备访问 Device Access：能够处理各种输入和输出设备



> 地理位置、陀螺仪、检测设备方向...



Tips：经常听到有人用“H5”去指代前端开发、前端项目，可以发现这样的叫法是不准确的，HTML5 本身表示 HTML 语言的一个版本、一系列功能的集合。当然为了和各个职能的同学方便对齐上下文，我们一般不特意指出这种认识上的“错误”。


**重点讲解**


#### WebSockets


- 解决了什么？
HTTP1.x 协议缺陷之一：通信只能由客户端发起，服务端不能主动给客户端推送，对于实时性场景不友好。
- 是什么？
是一种在客户端与服务器之间保持TCP长连接的网络协议，这样它们就可以随时进行信息交换。
虽然任何客户端或服务器上的应用都可以使用WebSocket，但原则上还是指浏览器与服务器之间使用。通过WebSocket，服务器可以直接向客户端发送数据，而无须客户端周期性的请求服务器，以动态更新数据内容。
- 优点是什么？
相对于短轮询、长轮询更为高效实时和资源有效性。
- 特点是什么？
没有跨域限制
- 应用？
弹幕
聊天室



#### 离线存储


##### 浏览器缓存


- 分类
强缓存和协商缓存：是否需要向服务器验证本地缓存是否依旧有效
- 强缓存控制字段：Cache-Control和Expire（HTTP1.0标准，可忽略，现在浏览器默认使用HTTP1.1）
Expire 控制缓存的原理：使用客户端的时间与服务端返回的时间做对比，那么如果客户端与服务端的时间因为某些原因（例如时区不同；客户端和服务端有一方的时间不准确）发生误差，那么强制缓存则会直接失效。
例如 设置expire的值为：Mon , 16 Apr 2020 01:12:33 GMT (绝对时间)
Cache-Control使用的是相对时间；例如设置Cache-Control的值为“public, max-age=xxx”，表示在xxx秒内再次访问该资源，均使用本地的缓存，不再向服务器发起请求。
   - public：所有内容都将被缓存（客户端和代理服务器都可缓存）
   - private：所有内容只有客户端可以缓存，Cache-Control的默认取值
   - no-cache：客户端缓存内容，但是是否使用缓存则需要经过协商缓存来验证决定
   - no-store：所有内容都不会被缓存，即不使用强制缓存，也不使用协商缓存


强缓存的缓存存放位置：

   - from memory cache (内容缓存)
   - from disk cach （硬盘缓存）
在浏览器中，浏览器会在js和图片等文件解析执行后直接存入内存缓存中，那么当刷新页面时只需直接从内存缓存中读取(from memory cache)；而css文件则会存入硬盘文件中，所以每次渲染页面都需要从硬盘读取缓存(from disk cache)


缺点：若后端接口更换，且缓存结果仍有效，则信息无法及时更新. 也可通过更改资源路径，强制刷新

- 协商缓存的字段：Last-Modified / If-Modified-Since   和   Etag / If-None-Match（后者优先级高）
区别：前者返回资源最后一次被修改时间 ； 后者返回资源唯一标识；
Last-Modified 的不足：
由于last-modified依赖的是保存的绝对时间，还是会出现误差的情况：
   1. 保存的时间是以秒为单位的，1秒内多次修改是无法捕捉到的；
   2. 各机器读取到的时间不一致，就有出现误差的可能性。为了改善这个问题，提出了使用E-tag。
- 强制缓存优先于协商缓存进行
强缓存与协商缓存需要配合起来使用才有意义.  当服务端希望客户端浏览器对某一资源进行缓存时，为了免去客户端每次都要询问自己：我上次的缓存现在还能用吗？所以，服务端选择了放权。只去告诉浏览器，我这次给你的资源你可以用多长时间，在这个时间段内，你可以一直使用它，无需每次咨询我。
强缓存是前端性能优化最有力的工具，对于有大量静态资源的网页，利用强缓存，提高响应速度。



##### 浏览器存储


###### Storage


- localStorage
只读的localStorage 属性允许你访问一个Document 源（origin）的对象 Storage；存储的数据将保存在浏览器会话中。
无论数据存储在 localStorage 还是 sessionStorage ，它们都特定于页面的协议。
另外，localStorage 中的键值对总是以字符串的形式存储。 (数值类型会自动转化为字符串类型).





- sessionStorage
允许你访问一个对应当前源的 session Storage 对象。它与 localStorage 相似，不同之处在于 localStorage 里面存储的数据没有过期时间设置，而存储在 sessionStorage 里面的数据在页面会话结束时会被清除。
页面会话在浏览器打开期间一直保持，并且重新加载或恢复页面仍会保持原来的页面会话。
- 区别与特点：
   - localStorage 的保存数据时间持久,除非被清理，否则一直存在。
   - 两者都保存在浏览器端；遵循同源策略。



###### IndexDB


IndexedDB 是一个运行在浏览器上的非关系型数据库。


#### History API


路由的概念来源于服务端，在服务端中路由描述的是 URL 与处理函数之间的映射关系。


- 前端路由
**前端路由**是现代SPA应用必备的功能,每个现代前端框架都有对应的实现,例如vue-router、react-router；路由描述的是 URL 与 UI 之间的映射关系，这种映射是单向的，即 URL 变化引起 UI 更新（无需刷新页面）
   - Hash 模式 ： Hash路由一个明显的标志是带有`#`,我们主要是通过监听url中的hash变化来进行路由跳转。
   - History模式
- 常用API
`history.back()`
`history.go()`
`history.forward`  相当于 `history.go(1)`> HTML5引入了 history.pushState() 和 history.replaceState() 方法，它们分别可以添加和修改历史记录条目。这些方法通常与window.onpopstate 配合使用。


`history.pushState()` 用于在浏览历史中添加历史记录,但是并不触发跳转,此方法接受三个参数，依次为：> `state`:一个与指定网址相关的状态对象，`popstate`事件触发时，该对象会传入回调函数。如果不需要这个对象，此处可以填`null`。
`title`：新页面的标题，但是所有浏览器目前都忽略这个值，因此这里可以填`null`。
`url`：新的网址，必须与当前页面处在同一个域。浏览器的地址栏将显示这个网址。


`history.replaceState()`  方法参数与前者一样，但是区别是修改浏览历史中当前纪录,而非添加记录,同样不触发跳转。



      `popstate`事件,每当同一个文档的浏览历史（即history对象）出现变化时，就会触发popstate事件。


        需要注意的是，仅仅调用`pushState`方法或`replaceState`方法 ，并不会触发该事件，只有用户点击浏览器     倒退按钮和前进按钮，或者使用 JavaScript 调用`back`、`forward`、`go`方法时才会触发。


         另外，该事件只针对同一个文档，如果浏览历史的切换，导致加载不同的文档，该事件也不会触发。


#### Web Worker


通过使用Web Workers，Web应用程序可以在独立于主线程的后台线程中，运行一个脚本操作。这样做的好处是可以在独立线程中执行费时的处理任务，从而允许主线程（通常是UI线程）不会因此被阻塞/放慢。


- 特点
高效
并行
- 适用
高密集大型计算
消耗主线程过大，如渲染
- 限制
无法访问DOM
- 基础API

```javascript
// src/main.js 主线程

const worker = new Worker("../src/worker.js"); 

//聆听信息
worker.onmessage = e => {
  const message = e.data;
  console.log(`[From Worker]: ${message}`);
  
  // 收到消息连续通信
  const reply = setTimeout(() => worker.postMessage("Marco!"), 3000);
};

// 发送信息
worker.postMessage("Marco!");


// src/worker.js

onmessage = e => {
  const message = e.data;
  console.log(`[From Main]: ${message}`);

  postMessage("Polo!");
};
```




#### requestAnimationFrame


- 几个概念
   - 视觉暂留：是光对视网膜所产生的视觉，在光停止作用后，仍然保留一段时间的现象，其具体应用是电影的拍摄和放映。原因是由视神经的反应速度造成的，其时值约是1/16秒，对于不同频率的光有不同的暂留时间。是动画、电影等视觉媒体形成和传播的根据。比如：我们日常使用的日光灯每秒大约熄灭100余次，但我们基本感觉不到日光灯的闪动。这都是因为视觉暂留的作用。所以，要达成最基本的视觉暂留效果至少需10fps（参考视频的帧率） 物体在快速运动时，当人眼所看到的影像消失后，人眼仍能继续保留其影像，约0.1-0.4秒左右的图像。
   - 动画的实现原理，是利用了人眼的“视觉暂留”现象，在短时间内连续播放数幅静止的画面，使肉眼因视觉残象产生错觉，而误以为画面在“动”。
   - 帧：在动画过程中，每一幅静止画面即为一“帧”。
   - 帧率：即每秒钟播放的静止画面的数量，单位是fps(Frame per second)。
   - 帧时长：即每一幅静止画面的停留时间，单位一般是ms(毫秒)。
- ##### 身边的帧率（频率）：

   - 10 FPS 达成基本视觉暂留
   - 25～30 FPS 传统广播电视信号
   - **60 FPS** 浏览器渲染刷新频率
   - 60～85 HZ 显示器刷新频率
   - 100 HZ 日光灯管闪烁频率
- 帧率反映动画流畅程度
   - 在网页中，帧率能够达到50~60fps的动画将会相当流畅，让人倍感舒适。
   - 帧率在30～50fps之间的动画，因各人敏感程度不同，舒适度因人而异。
   - 帧率在30fps以下的动画，让人感觉到明显的卡顿和不适感。
   - 帧率波动很大的动画，亦会使人感觉到卡顿。


流畅的动画一般具备两个特点：

   - **帧率高**（接近60fps最佳）
   - **帧率稳定，波动少**（极少出现跳帧现象）


为什么电影/电视的帧率小于30fps，但依然感觉很流畅？ 这个问题可以参阅[这条知乎回答](https://www.zhihu.com/question/21081976#answer-9265972)

- 失帧
首先，浏览器会以最大m次/秒刷新屏幕。数字m取决于电脑的屏幕刷新率，浏览器的刷新率，以及CPU、GPU的处理能力。如果你的浏览器只能以30帧/s的速度刷新屏幕（由于上面的一个或者多个原因造成），那么以60帧/秒的速度运行动画是没有什么意义的，多余的帧数将会消失。与此同时，对DOM结构所做的更改要比浏览器渲染的要多，这也被称为布局抖动，因为这些操作是同步的，会影响网站的性能以及绘制操作，从而导致动画效果不佳。
此时，需要来自浏览器的某种回调函数，他会告诉我们下一次屏幕刷新的时间，或者更准确的说，是下一次绘制操作将在何时执行。这个回调函数就是requestAnimationFrame Web API。
- 背景
传统的 `javascript` 动画是通过定时器 `setTimeout` 或者 `setInterval` 实现的。但是定时器动画一直存在两个问题，第一个就是动画的循时间环间隔不好确定，设置长了动画显得不够平滑流畅，设置短了浏览器的重绘频率会达到瓶颈，推荐的最佳循环间隔是17ms（大多数电脑的显示器刷新频率是60Hz，1000ms/60）；第二个问题是定时器第二个时间参数只是指定了多久后将动画任务添加到浏览器的UI线程队列中，如果UI线程处于忙碌状态，那么动画不会立刻执行。为了解决这些问题，H5 中加入了 `requestAnimationFrame`;
- 优点
   1. `requestAnimationFrame` 会把每一帧中的所有 DOM 操作集中起来，在一次重绘或回流中就完成，并且重绘或回流的时间间隔紧紧跟随浏览器的刷新频率
   2. 在隐藏或不可见的元素中，`requestAnimationFrame` 将不会进行重绘或回流，这当然就意味着更少的 CPU、GPU 和内存使用量
   3. `requestAnimationFrame` 是由浏览器专门为动画提供的 API，在运行时浏览器会自动优化方法的调用，并且如果页面不是激活状态下的话，动画会自动暂停，有效节省了 CPU 开销。
- `requestAnimationFrame` 动画

```javascript
const deg = 0;
const id;
const div = document.getElementById("div");
div.addEventListener('click', function () {
    var self = this;
    requestAnimationFrame(function change() {
        self.style.transform = 'rotate(' + (deg++) + 'deg)';
        id = requestAnimationFrame(change);
    });
});
document.getElementById('stop').onclick = function () {
    cancelAnimationFrame(id);
};
```

- 大数据渲染

```javascript
const total = 100000;
const size = 100;
const count = total / size;
const done = 0;
const ul = document.getElementById('list');

function addItems() {
    var li = null;
    var fg = document.createDocumentFragment();

    for (var i = 0; i < size; i++) {
        li = document.createElement('li');
        li.innerText = 'item ' + (done * size + i);
        fg.appendChild(li);
    }

    ul.appendChild(fg);
    done++;

    if (done < count) {
        requestAnimationFrame(addItems);
    }
};

requestAnimationFrame(addItems);
```

性能会比定时器提升很多。



#### Canvas


提供了一个空白绘图区域，可以使用 APIs （比如 Canvas 2D 或 WebGL）来绘制图形。


- 基础API ( 参考MDN)

```javascript
// 画一个绿色的矩形
ctx.beginPath(); // 开始路径绘制
ctx.moveTo(20, 20); // 设置路径起点，坐标为(20,20)
ctx.lineTo(200, 20); // 绘制一条到(200,20)的直线
ctx.lineWidth = 1.0; // 设置线宽
ctx.strokeStyle = "#CC0000"; // 设置线的颜色
ctx.stroke(); // 进行线的着色，这时整条线才变得可见

// 画一个粉色的线
ctx.fillStyle = "green";
ctx.fillRect(10, 10, 150, 100);

ctx.strokeRect(10, 10, 200, 100);

// 设置字体
ctx.font = "Bold 20px Arial";
// 设置对齐方式
ctx.textAlign = "left";
// 设置填充颜色
ctx.fillStyle = "#008600";
// 设置字体内容，以及在画布上的位置
ctx.fillText("Hello!", 10, 50);
// 绘制空心字
ctx.strokeText("Hello!", 10, 100);
```

- 应用
   - 截屏 [参考插件 html2canvas](http://html2canvas.hertzen.com/)
   - 图形验证码
   - 画板



### 作业


1. 用 Vue2 写一个 todolist，通过提供的 api：[jsonplaceholder](https://jsonplaceholder.typicode.com/) 的 todos api 获取初始的 todolist，之后删除完成的请求不作要求（Vue2 应用的创建可以直接用 [vue-cli](https://cli.vuejs.org/zh/)，当然能用 Webpack 自己搭建就更好了）
![](https://ahabhgk.github.io/static/4ea87b07ae61e0aa6cf5d79cccef7bd9/7d769/todos-api.png#align=left&display=inline&height=301&margin=%5Bobject%20Object%5D&originHeight=301&originWidth=960&status=done&style=none&width=960)
2. 现在 [Slowly Render](https://github.com/ahabhgk/slowly-render) 已经有了 HashRouter，为了更完善的使用方式，用 history api 完成 HistoryRouter 吧（要求使用方式同 HashRouter）
3. 将作业 1 用 Vue3 的 [Composition API](https://composition-api.vuejs.org/zh/api.html) 进行重构，并阅读 [Composition API 的 RFC](https://composition-api.vuejs.org/zh/)
Vue3 尝鲜方式：
   - 直接用 [vite 创建 Vue3 应用](https://github.com/vitejs/vite)
   - 在 vue-cli 中使用 [vue-cli-plugin-vue-next](https://github.com/vuejs/vue-cli-plugin-vue-next)
   - 在 Vue2 中使用 [@vue/composition-api 插件](https://github.com/vuejs/composition-api/blob/master/README.zh-CN.md)



### REFs


[一名【合格】前端工程师自检清单](https://juejin.im/post/5cc1da82f265da036023b628#heading-36)


[前端存储最佳实践](https://juejin.im/post/5c136bd16fb9a049d37efc47)


[精读《谈谈Web Workers》](https://juejin.im/post/5bcd1e895188256e226580da)


[Web 动画性能指南](https://alexorz.github.io/animation-performance-guide/)


[requestAnimationFrame 理解与实战](https://newbyvector.github.io/2018/05/01/2015-05-01/)
