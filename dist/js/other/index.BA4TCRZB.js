import{b as u,o as s,c as p,ad as k,w as n,e as o,g as a,al as z,V as x,ae as S,ah as v,ag as B,am as C,an as N,ao as w,ap as j,aq as T,$ as y,q as b,Z as $,t as P,f as m,ar as V,as as H,at as X,ai as q,h,i as D}from"../.pnpm/.pnpm.Bbu8tIZm.js";const I=u({__name:"index",props:{tooltipConfig:{type:Object,default:()=>{}},popConfig:{type:Object,default:()=>{}},popText:{type:String,default:""},tooltipText:{type:String,default:""},icon:{type:String,default:""},iconSize:{type:[Number,String],default:24}},setup(e){return(r,_)=>{const t=w,c=j,f=T;return s(),p(f,C(N(e.popConfig)),k({trigger:n(()=>[o(c,x({trigger:"hover"},e.tooltipConfig),{trigger:n(()=>[S(r.$slots,"default")]),default:n(()=>[a("span",null,v(e.tooltipText),1)]),_:3},16)]),default:n(()=>[B(" "+v(e.popText),1)]),_:2},[e.icon?{name:"icon",fn:n(()=>[o(t,{size:e.iconSize},{default:n(()=>[a("span",{class:z(e.icon)},null,2)]),_:1},8,["size"])]),key:"0"}:void 0]),1040)}}}),O=u({__name:"index",props:{pane:{type:Boolean,default:!1},tablist:{type:Array,default:()=>[]}},setup(e){const r=t=>u({render(){return typeof t=="function"?t():t}}),_=t=>{const c={...t};return delete c.children,c};return(t,c)=>{const f=V,d=H,g=X;return s(),p(g,x({type:"line",animated:""},t.$attrs),k({default:n(()=>[e.pane?(s(!0),b($,{key:0},y(e.tablist,(i,l)=>(s(),p(f,x({key:l,ref_for:!0},_(i)),{default:n(()=>[i.children?(s(),p(P(r(i.children)),{key:0})):m("",!0)]),_:2},1040))),128)):m("",!0),e.pane?m("",!0):(s(!0),b($,{key:1},y(e.tablist,(i,l)=>(s(),p(d,{key:l,name:i},null,8,["name"]))),128))]),_:2},[y(t.$slots,(i,l)=>({name:l,fn:n(()=>[S(t.$slots,l,{slot:i})])}))]),1040)}}}),R={key:0},A=u({__name:"index",props:{text:{type:String,default:""},icon:{type:String,default:""},iconSize:{type:[Number,String],default:24}},setup(e){return(r,_)=>{const t=w,c=q;return s(),p(c,C(N(r.$attrs)),k({default:n(()=>[e.text?(s(),b("span",R,v(e.text),1)):m("",!0)]),_:2},[e.icon?{name:"icon",fn:n(()=>[o(t,{size:e.iconSize},{default:n(()=>[a("span",{class:z(e.icon)},null,2)]),_:1},8,["size"])]),key:"0"}:void 0,e.text?void 0:{name:"default",fn:n(()=>[S(r.$slots,"default")]),key:"1"}]),1040)}}}),E={class:"p-4"},F=a("div",{class:"border-0 border-l-4 border-solid border-green-400 text-2xl font-bold pl-4"},"按钮用法",-1),L={class:"py-10"},Z=a("div",{class:"border-0 border-l-4 border-solid border-green-400 text-2xl font-bold pl-4"},"tabs示例(默认不是使用面板)",-1),G={class:"py-10"},J=a("div",{class:"border-0 border-l-4 border-solid border-green-400 text-2xl font-bold pl-4"},"tabs示例使用面板",-1),K={class:"py-10"},M=a("div",{class:"border-0 border-l-4 border-solid border-green-400 text-2xl font-bold pl-4"},"信息弹窗提示",-1),Q={class:"py-10"},W=u({__name:"index",setup(e){const r=[{name:"tabs1",children:h("div",{},"Hello1")},{name:"tabs2",children:h("div",{},"Hello2")},{name:"tabs3",children:h("div",{},"Hello3")}],_=()=>{window.$msg.info("是的")},t=()=>{window.$msg.info("并不")};return(c,f)=>{const d=A,g=D,i=O,l=I;return s(),b("div",E,[F,a("div",L,[o(g,null,{default:n(()=>[o(d,{type:"info",iconSize:30,size:"large",text:"按钮111",icon:"iconify-color emojione-v1--baby-chick"}),o(d,{type:"error",text:"error"}),o(d,{tertiary:"",circle:"",iconSize:"30",icon:"iconify-color emojione-v1--bear-face",size:"large"})]),_:1})]),Z,a("div",G,[o(i,{tablist:["tab1","tab2","tab3"]})]),J,a("div",K,[o(i,{pane:!0,tablist:r})]),M,a("div",Q,[o(l,{icon:"iconify-color emojione-v1--airplane","pop-config":{onPositiveClick:_,onNegativeClick:t},"tooltip-config":{placement:"bottom"},"pop-text":"确认删除吗","tooltip-text":"删除"},{default:n(()=>[o(d,{type:"error",text:"error"})]),_:1},8,["pop-config"])])])}}});export{W as default};