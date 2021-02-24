
1. 由于在执行setup函数时候，还没有执行created生命周期方法,所以在setup函数中，是无法使用data和methods

2. 由于我们不能在setup函数中使用data和methods,所以VUE为了避免我们错误的使用,它直接将setup函数中this修改成了undefined

3. setup函数只能是同步的不能是异步的

https://www.cnblogs.com/rmlzy/p/14343325.html
```
export {
  ref, // 代理基本类型
  shallowRef, // ref 的浅代理模式
  isRef, // 判断一个值是否是 ref
  toRef, // 把响应式对象的某个 key 转为 ref
  toRefs, // 把响应式对象的所有 key 转为 ref
  unref, // 返回 ref.value 属性
  proxyRefs,
  customRef, // 自行实现 ref						
  triggerRef, // 触发 customRef
  Ref, // 类型声明
  ToRefs, // 类型声明
  UnwrapRef, // 类型声明
  ShallowUnwrapRef, // 类型声明
  RefUnwrapBailTypes // 类型声明
} from './ref'
export {
  reactive, // 生成响应式对象
  readonly, // 生成只读对象
  isReactive, // 判断值是否是响应式对象
  isReadonly, // 判断值是否是只读对象
  isProxy, // 判断值是否是 proxy
  shallowReactive, // 生成浅响应式对象
  shallowReadonly, // 生成浅只读对象
  markRaw, // 让数据不可被代理
  toRaw, // 获取代理对象的原始对象
  ReactiveFlags, // 类型声明
  DeepReadonly // 类型声明
} from './reactive'
export {
  computed, // 计算属性
  ComputedRef, // 类型声明
  WritableComputedRef, // 类型声明
  WritableComputedOptions, // 类型声明
  ComputedGetter, // 类型声明
  ComputedSetter // 类型声明
} from './computed'
export {
  effect, // 定义副作用函数, 返回 effect 本身, 称为 runner
  stop, // 停止 runner
  track, // 收集 effect 到 Vue3 内部的 targetMap 变量
  trigger, // 执行 targetMap 变量存储的 effects
  enableTracking, // 开始依赖收集
  pauseTracking, // 停止依赖收集
  resetTracking, // 重置依赖收集状态
  ITERATE_KEY, // 固定参数
  ReactiveEffect, // 类型声明
  ReactiveEffectOptions, // 类型声明
  DebuggerEvent // 类型声明
} from './effect'
export {
  TrackOpTypes, // track 方法的 type 参数的枚举值
  TriggerOpTypes // trigger 方法的 type 参数的枚举值
} from './operations'
```
名词解释
```
target: 普通的 JS 对象

reactive: @vue/reactivity 提供的函数, 接收一个对象, 并返回一个 代理对象, 即响应式对象

shallowReactive: @vue/reactivity 提供的函数, 用来定义浅响应对象

readonly:@vue/reactivity 提供的函数, 用来定义只读对象

shallowReadonly: @vue/reactivity 提供的函数, 用来定义浅只读对象

handlers: Proxy 对象暴露的钩子函数, 有 get()、set()、deleteProperty()、ownKeys() 等, 可以参考MDN

targetMap: @vue/reactivity 内部变量, 存储了所有依赖

effect: @vue/reactivit 提供的函数, 用于定义副作用, effect(fn, options) 的参数就是副作用函数

watchEffect: @vue/runtime-core 提供的函数, 基于 effect 实现

track: @vue/reactivity 内部函数, 用于收集依赖

trigger: @vue/reactivity 内部函数, 用于消费依赖

scheduler: effect 的调度器, 允许用户自行实现
```

### setup中的data如何使用

```
import { ref, toRefs, reactive }  from 'vue'
export default {
  name: 'Test',
  setup(){
    // 在setup中定义template模板使用的响应式变量，你得用ref或者reactive来定义
    // 这里的ref你可以理解成一个工厂类，传入的参数就是初始化的变量的值，跟组件的ref概念不能混淆
    // 定义单个变量，为了让大家明白意义，我跟vue2.0都进行下对比 
    // vue2.0,不管定义单个或者多个我们都是定义 在data里，如
    /*
       data(){
          return {
            name: '小哥哥'
          }
       }
    */
    // 在vue3.0中，如果你只需要定义一个响应式变量，那么你可以用以下ref
    // 可能你会疑惑既然是定义变量为什么不用let，var，而用const定义常量的，这里是因为你定义的是一个引用，name指向的永远是一个固定不变的指针地址
    const name = ref('小哥哥')
    // 注意点，这里要获取name的值怎么获取,通过定义的变量的。value
    console.log('拿到name的值：', name.value)
    // 检测某个值是不是响应式的可以用isRef
    
    // 每次都用.value去拿值的写法，是不是有点不适应，而且定义的变量多了我们也不可能定义一堆ref，看起来都丑
   // reactive 可以创建一个响应式数据，参数是一个对象
    const data = reactive({
       name: '帅的没人要'，// 这样创建的响应式数据就不用ref了，写起来是不是要方便点
        sex: '性别分不出'，
        arr: []
    })
    // 但是以上还是有个缺点，如果你在return里直接放上data,那你使用的时候每次都要data.name，data.sex也麻烦，<div>{{data.sex}}</div>
   // 你说你可以解构在return {...data} 如果你这样的，里面的数据都会变成非响应式的，vue3.0提供了一个满足你幻想的方法toRefs,使用了这个包装下，你就可以在return里使用解构了
 // 包装上面的data
  const  refData = toRefs(data)
    // 在data中都有个return ，这里当然也必须要有的，但是这里是所有方法计算属性都要通过这个返回
    // 有人疑惑，那我可以直接在这个return上定义变量吗，答案是可以，但是定义的变量不是响应式的，不可变化
    return {
      ...refData, // 你也可以直接在这里用...toRefs(data) 这样会简单点
      name,
      rules: [] //如果你有什么不会变化的数据，如规则啊，配置啊，可以直接在这定义而不需要通过ref和reactive
    }
    
  }
}
```

### setup中的method如何使用
```
import { ref } from 'vue'
export default {
  name: 'Test',
  setup(){
    // 定义一个响应式数据
    const baby = ref('1岁bb')
   // 定义method,把方法名字在return里返回
   // 注意：这里用调用响应式的数据也就是您定义的vue2.0的data,不可以用this,这个setup函数在01里已经说明过了，这个时候相当于vue2.0的beforeCreate和created，这里一开始就会运行，还没有this的指向，是个undefined，访问所有你定义的响应式的变量都要.value去访问
    const wantToKnowBabysName = () => {
      console.log("oh,he is a " + baby.value)
    }
    const getData = () => {}
   // 对比vue2.0
   /*
   method: {
      wantToKnowBabysName(){
        console.log("oh,he is a " + baby.value)
      },
      getData() {
      }
    }
  */
   
    return {
      baby,
      wantToKnowBabysName,
      getData
    }
    
  }
}
```

### setup中的computed如何使用
```
// 注意使用的时候引入computed
import { ref, computed} from 'vue'
export default {
  name: 'Test',
  setup(){
    // 定义一个响应式数据
    const baby = ref('嘎嘎嘎')
    // 定义翠花年龄
    const age = ref(28)
    // computed计算属性传入一个回调函数
    const areYouSureYouAreABaby = computed(() => {
      return `I'm sure,I'm a 300 month baby, ${baby.value}`
    })
    // set和get方式
    const setAge= computed({
      get() {
        return age.value + 10
      },
      set(v) {
        age.value = v - 10
      }
    })
   // 对比vue2.0
   /*
   computed: {
      areYouSureYouAreABaby (){
        return `I'm sure,I'm a 300 month baby, ${baby.value}`
      },
      setAge:{
        get(){
          return age + 10
        },
        set(v) {
          age = v - 10
        }
      }
    }
  */
   
    return {
      baby,
      age,
      areYouSureYouAreABaby 
    }
    
  }
}
```

### setup中的watch如何使用
```
// 注意使用的时候引入watch
import { ref, watch, watchEffect } from 'vue'
export default {
  name: 'Test',
  setup(){
    // 定义一个响应式数据
    const baby = ref('嘎嘎嘎')
    const arr = ref(['翠花', '小红'])
    // 监听一个值的情况，有两种方式
    // watch 有三个参数：第一个是个getter（所谓getter写法就是你要写个getter函数）,第二个是个回调函数，第三个是个options(这个参数是放vue2.0的deep或者immediate等可选项)
    // 第一种：直接放ref
    watch(baby, () => {
      return `I'm sure,I'm a 300 month baby, ${baby.value}`
    })
   // 第二种：放ref的value值
   watch(() => baby.value, () => {
      return `I'm sure,I'm a 300 month baby, ${baby.value}`
    })
  
   // 监听多个值的时候 ,第一个参数是个数组，里面放监听的元素
   watch([baby,arr], (v, o) => { 
     // 这里的v,o也是数组，所以你取值的时候v[0],v[1]拿到第几个元素的变化
     ...
   }, {
    deep: true,
    immediate: true
   })
 // 或者写成
 watch([baby,arr], ([baby, arr], [prebaby,prearr]) => {
    ...
  })
   // 对比vue2.0
   /*
   watch: {
      baby(v, o) {
        
      },
      arr: {
        handle(v,o) {},
        deep: true,
        immediate: true,
        flush: 'pre' // 这个默认有三个参数，'pre'| 'post' | 'sync'，默认‘pre’组件更新前运行,'post'组件渲染完毕后执行，一般用于你需要去访问$ref的时候可以用这个，'sync'是一旦你的值改变你需要同步执行回调的时候用这个
      }
    }
  */
   // watch的写法跟vue2的$watch写法一样，可以参考
  // vue3.0 watchEffect 用法
  //  watchEffect 有两个参数，一个是副作用函数(就是外部的数据对这个函数产生影响的，通俗点说就是在这个函数内部使用了外面的变量等)，一个是options（）
//  在vue2.0中，我们一般在created里添加一些监听事件，比如你的$bus的一些事件监听，在setup中就可以在这个里面写
watchEffect((onInvalidate) => {
   // 这里的变化就相当于依赖了age.value，如果age变化了就会触发这个监听
   // 刚刚创建组件的时候会立即执行这个 
   const _age= `her age is ${age.value}`
   console.log(_age)
   //有时候你需要在这里挂载一些监听事件
   const handerClick = ()=>{}
   document.addEventlistener('click', handerClick)
   // 在vue2.0我们需要在destroy的时候remove它，这里提供了一个方法onInvalidate回调解决remove的问题
   onInvalidate(()=>{
       /*
        执行时机:  在副作用即将重新执行时，就是在每次执行这个watchEffect回调的时候会先执行这个,如果在setup()或生命周期钩子函数中使用watchEffect, 则在卸载组件时执行此函数。
       */
       document.removeEventListener('click',handerClick )
    })  
})
// 这个也是支持async,await的
const data = ref(null)
watchEffect(async onInvalidate => {
 // 假设个接口获取数据的
  data.value = await fetchData()
  onInvalidate(() => {...})
})
// 再来理解options：这里有三个参数flush,onTrigger,onTrack
watchEffect(onInvalidate => {
  onInvalidate(() => {...})
}, {
  flush: 'pre' // 跟watch一样，默认pre，组件更新前去调用
  onTrigger(e) {
    // 依赖项变化时候触发这个即依赖项的set触发的时候
  },
  onTrack(e) {
    // 依赖项被调用的时候触发这个即依赖项的get触发的时候
  }
})
    return {
      baby,
      areYouSureYouAreABaby,
      data 
    }
    
  }
}
```

### props在setup中的使用
setup接受两个参数props和context
```
export default {
  name: 'Test',
  props: ['name', 'age'],
  // setup(props, context) { // 有的时候会这样写，你可能只用得到emit
  setup(props,{attrs, slots, emit}) // 如果你都用得到你可以这样解构的写出来，这个不是响应式的，所以可解构
   // 错误写法 const {name} = props 这里我理解你肯定想直接就使用name，age等
   // 这个props是一个响应式的Proxy对象，不可以解构，解构后悔失去响应，如果要用解构的方式，要用toRefs
    let { name, age } = toRefs(props)
    // 现在是不是感觉可以直接就用操作name和age了，天真，转成ref了，所有的访问值都要xx.value
   console.log(name.value,age.value)
   // 所以倒回去，是不是觉得还不如直接用props.name直接访问代理对象的值要好点
   console.log(props.name, props.age)

   // context 
   // 看到这个context的参数你应该知其意了撒
   //  attrs: 相当于vue2.0的$attrs,意思就是传进来的属性值除了props接受的那部分
   // slots: 就是插槽，你要获取插槽的什么值的话可以用这个
   // 插槽这里解释下，可能有部分人不知道咋拿值，顺便说下，这里vue3.0的所有响应式数据都是Proxy对象，所以你取值的时候都是proxy.xxx获取
   slots.default(args => {
    console.log('这里就是你在vue2.0里看到的所有slot的数组，就可以取你要哪个插槽的值了', args)
   })
   // emit: 这个是vue2.0 的$emit
   emit('方法名', '参数') // vue2.0 this.$emit('方法名', '参数')
    ...
  }
}
```

### provide和inject
父组件
```
// 老父亲组件
import { provide, ref, reactive } from 'vue'
export default {
  name: 'Test',
  setup() {
    // 用法: provide(key, value) 用下面的ref和reactive是为了让数据变成响应式的，父组件变化，子组件数据跟着变
    const name = ref('小哥哥')
    const obj = reactive({
      name: '土狗',
      age: '3岁'
    })
    provide('name', name)
    provide('animal', obj)
  }
}
```
子组件
```
// 乖儿孙组件
import { inject } from 'vue'
export default {
  name: 'Child,
  setup() {
    // 用法: inject(key) 
    const name = inject('name')
    const animal = inject('animal')
    return {
      name,
      animal
    }
  }
}
```

### Composition API
Vue2 的 Options API 和 Composition API 的区别：
```
/* Options API */
export default {
  props: {},
  data(){},
  computed: {},
  watch: {},
  methods: {},
  created(),
  components:{}
  // ...other options
}

/* Composition API */
export default {
  props: {},
  setup(),
  components:{}
}
```
这就是两种 API 在大致结构上的不同，虽然 Composition API 提倡使用 setup 来暴露组件的 data、computed、watch、生命周期钩子... 但并不意味着强制使用，在 Vue3 中同样可以选择 Options API 或者两种写法混用。
#### setup
setup 方法接受两个参数 setup(props, context) ，props 是父组件传给组件的数据，context(上下文) 中包含了一些常用属性：

#### context.attrs 表示由上级传向该组件，但并不包含在 props 内的属性：

```
<!-- parent.vue -->
<Child msg="hello world" :name="'child'"></Child>
/* child.vue */
export default {
  props: { name: String },
  setup(props, context) {
    console.log(props) // {name: 'child'}
    console.log(context.attrs) // {msg: 'hello world'}
  },
}
```
#### context.emit用于在子组件内触发父组件的方法
```
<!-- parent.vue -->
<Child @sayWhat="sayWhat"></Child>
/* child.vue */
export default {
  setup(_, context) {
    context.emit('sayWhat')
  },
}
```
#### slots

用来访问被插槽分发的内容，相当于 vm.$slots
```
<!-- parent.vue -->
<Child>
  <template v-slot:header>
    <div>header</div>
  </template>
  <template v-slot:content>
    <div>content</div>
  </template>
  <template v-slot:footer>
    <div>footer</div>
  </template>
</Child>
/* child.vue */
import { h } from 'vue'
export default {
  setup(_, context) {
    const { header, content, footer } = context.slots
    return () => h('div', [h('header', header()), h('div', content()), h('footer', footer())])
  },
}
```

### 生命周期
Vue3 的生命周期除了可以使用传统的 Options API 形式外，也可以在 setup 中进行定义，只不过要在前面加上 on：
```
export default {
  setup() {
    onBeforeMount(() => {
      console.log('实例创建完成，即将挂载')
    })
    onMounted(() => {
      console.log('实例挂载完成')
    })
    onBeforeUpdate(() => {
      console.log('组件dom即将更新')
    })
    onUpdated(() => {
      console.log('组件dom已经更新完毕')
    })
    // 对应vue2 beforeDestroy
    onBeforeUnmount(() => {
      console.log('实例即将解除挂载')
    })
    // 对应vue2 destroyed
    onUnmounted(() => {
      console.log('实例已经解除挂载')
    })
    onErrorCaptured(() => {
      console.log('捕获到一个子孙组件的错误')
    })
    onActivated(() => {
      console.log('被keep-alive缓存的组件激活')
    })
    onDeactivated(() => {
      console.log('被keep-alive缓存的组件停用')
    })
    // 两个新钩子，可以精确地追踪到一个组件发生重渲染的触发时机和完成时机及其原因
    onRenderTracked(() => {
      console.log('跟踪虚拟dom重新渲染时')
    })
    onRenderTriggered(() => {
      console.log('当虚拟dom被触发重新渲染时')
    })
  },
}
```
Vue3 没有提供单独的 onBeforeCreate 和 onCreated 方法，因为 setup 本身是在这两个生命周期之前执行的，Vue3 建议我们直接在 setup 中编写这两个生命周期中的代码。

### Reactive API
https://www.cnblogs.com/FrankLongger/p/14439342.html

### computed