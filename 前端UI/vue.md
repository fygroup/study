### 实例
应用中，只会创建一个Vue根实例，应用都是通过这个根实例启动的
```
import Vue from 'vue';
import App from './App.vue'
...
// 激活Vue调试工具vue-devtools
Vue.config.devtools = true;
new Vue({
    router, // react-router
    store,  // vuex
    el: '#app', // 需要渲染的DOM节点, 将其挂载到#app上
    render: h => h(App) // //h是createElement 的别名，创建/加载组件
});

//app.vue
<template>
  <div id="app">
    <router-view/>
  </div>
</template>
<script>
export default {
  name: 'App'
}
</script>
```

---
### 组件
```
(1)局部组件
局部组件只能在所定义的Vue实例中使用，格式如下：
//定义<my-component>组件
new Vue({
  // ...
  components: {
    // <my-component> 将只在父模板可用
    'my-component': {
      template: '<div>A custom component!</div>'
    }
  }
})

(2)全局组件
方法一：Vue.component('my-component',{})
方法二：my-component.vue


```

---
### 文本
```
<span>Message: {{ msg }}</span>
<span v-once>这个将不会改变: {{ msg }}</span>
```

---
### Vue的渲染机制
```
el ---> template ---> render ---> vnode 

//独立构建：包含模板编译器
渲染过程: html字符串 → render函数 → vnode → 真实dom节点

//运行时构建： 不包含模板编译器
渲染过程: render函数 → vnode → 真实dom节点

```


### Vue.extend、Vue.component与new Vue(vue构造、vue组件和vue实例)
```
关系：vue构造->vue组件->vue实例

<div id="mount-point"></div>
// 创建构造器
var Profile = Vue.extend({
  template: '<p>{{firstName}} {{lastName}} aka {{alias}}</p>',
  data: function () {
    return {
      firstName: 'Walter',
      lastName: 'White',
      alias: 'Heisenberg'
    }
  }
})
// 创建 Profile 实例，并挂载到一个元素上。
new Profile().$mount('#mount-point')


// 注册组件，传入一个扩展过的构造器
Vue.component('my-component', Vue.extend({ /* ... */ }))

// 注册组件，传入一个选项对象 (自动调用 Vue.extend)
Vue.component('my-component', { /* ... */ })

// 获取注册的组件 (始终返回构造器)
var MyComponent = Vue.component('my-component')


```

---
### extends和mixins
```
都有扩展组件的作用
extends对象跟mixins对象很类似、
extends只支持一个对象，而mixins支持的是数组
extends:countConsole,
mixins:[countConsole]

export default {
    extends: 
    mixins: []
}
```

---
### new vue中的el
```
<div id='id'></div>

new vue({
    el: '#id',   //新建的vue要挂载的DOM id
    render: (h)=>{
        h()
    }
})
```

---
### $mount
```
new Vue.extend({
    template: "<p></p>",
    //或者
    render: (h) => {
        return h('div')
    },

    data: function(){
        return {
            name:''
        }
    }
})().$mount('#id')


```

---
### render
```

```

---
### 小心store-getter
```
store的getter，不能进行watch！！！
必须将其赋值，然后对赋值进行监听
```

---
### v-if
```
不存在:v-if的用法
```

---
### props watch mounted
```
初始创建的组件并不会watch props，只有已经创建好的组件才会watch props


监听props要谨慎，如果监听的参数需要驱动一些东西的话，最好放在mounted中。因为组件的v-if创建，并不会刷新props，不会被watch监听，除非此组建已存在


初始的props不会被watch监听

```

---
### input 修改值
```
<Input v-model="value" @on-change="change"/>

data(){
  return {
    value: ''
  }
},
methods:{
  change(x){
    this.$nextTick(()=>{
      this.value = 'xxxxx'
    })
  }
}



```

### icon原始点击
```
<Icon @click.native="collapsedSider" .../>
```

### drawer在容器中展示
```
注意：
position:relative 表示在当前的div里面
如果不加就会在上一级的容器里面


<div style="position:relative;width: 800px;height: 600px;border: 1px solid pink;">
<div style="position:relative;width: 500px;height: 400px;border: 1px solid pink;">
    <Button @click="value1 = true" type="primary">Open</Button>
    <Drawer title="Basic Drawer" :closable="false" v-model="value1" inner :transfer="false">
        <p>Some contents...</p>
        <p>Some contents...</p>
        <p>Some contents...</p>
    </Drawer>
</div>  
</div>  
```

### iview table className
```
https://www.jianshu.com/p/4c8028a198d6

<style>
/*外层table的border*/
  .ivu-table-wrapper {
    border:none
  }
/*底色*/
.ivu-table td{
  background-color: #182328;
  color: #fff;
}
/*每行的基本样式*/
  .ivu-table-row td {
    color: #fff;
    border:none
  }
  /*头部th*/
  .ivu-table-header th{
    color:#FFD3B4;
    font-weight: bold;
    background-color: #212c31;
    border: none;
  }

  /*偶数行*/
  .ivu-table-stripe-even td{
    background-color: #434343!important;
  }
  /*奇数行*/
  .ivu-table-stripe-odd td{
    background-color: #282828!important;
  }
/*选中某一行高亮*/
  .ivu-table-row-highlight td {
    background-color: #d63333!important;
  }
  /*浮在某行*/
  .ivu-table-row-hover td {
    background-color: #d63333!important;
  }
</style>
<template>
  <div>
    <Table ref="selection" @on-selection-change="onSelect" height="700" no-data-text="暂无数据" :row-class-name="rowClassName" :columns="columns4" :data="data1" highlight-row></Table>
    <Button @click="handleSelectAll(true)">Set all selected</Button>
    <Button @click="handleSelectAll(false)">Cancel all selected</Button>
  </div>
</template>

```

### style scoped
```
最好不要加 scoped
```

### store的坑
```
store是个持久化的东西，只要没刷新界面，他就会一直存在。
如果涉及修改store.state的内容，那么在view层拷贝一份，然后再进行修改
```

### route与持久化
```
<keep-alive>
  <router-view/>
</keep-alive>

[
  {
    path: 'diseaseTaskList',
    name: 'diseaseTaskList',
    component: () => import('@/view/task/diseaseList'),
  }
]

在@/view/task/diseaseList/index.js文件中
  import diseaseTaskList from './diseaseTaskList.vue'
  export default diseaseTaskList

export的模块名必须与route中的path、name一样，否则不会持久化



```

### $listeners
```
//获得当前组件绑定的事件
this.$listeners
```

### form渲染
```
// 一定要先有form里的内容，再有渲染form的dom

//例如

<form>
  <formItem v-for="formDom">
    ...
  </formItem>
</form>

init (){
  form = {....}
  this.nextTick(()=>{
    formDom= {...}
  })
}

data () {
  return {
    form: {}
  }
}
form
```