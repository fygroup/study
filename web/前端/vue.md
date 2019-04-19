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