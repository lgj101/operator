"use strict";(self.webpackChunkweb_app=self.webpackChunkweb_app||[]).push([[112],{9505:(e,n,t)=>{t.d(n,{Z:()=>c});var a=t(72791),s=t(81207);const c=(e,n)=>{const[t,c]=(0,a.useState)(!1);return[t,(t,a,o,r)=>{c(!0),s.Z.invoke(t,a,o,r).then((n=>{c(!1),e(n)})).catch((e=>{c(!1),n(e)}))}]}},32112:(e,n,t)=>{t.r(n),t.d(n,{default:()=>m});var a=t(72791),s=t(51691),c=t(21435),o=t(61889),r=t(9505),i=t(40306),l=t(75952),p=t(87995),d=t(41320),h=t(80184);const m=e=>{let{deleteOpen:n,selectedPVC:t,closeDeleteModalAndRefresh:m}=e;const u=(0,d.TL)(),[C,x]=(0,a.useState)(""),[b,f]=(0,r.Z)((()=>m(!0)),(e=>u((0,p.Ih)(e))));return(0,h.jsx)(i.Z,{title:"Delete PVC",confirmText:"Delete",isOpen:n,titleIcon:(0,h.jsx)(l.NvT,{}),isLoading:b,onConfirm:()=>{C===t.name?f("DELETE","/api/v1/namespaces/".concat(t.namespace,"/tenants/").concat(t.tenant,"/pvc/").concat(t.name)):u((0,p.Ih)({errorMessage:"PVC name is incorrect",detailedError:""}))},onClose:()=>m(!1),confirmButtonProps:{disabled:C!==t.name||b},confirmationContent:(0,h.jsxs)(s.Z,{children:["To continue please type ",(0,h.jsx)("b",{children:t.name})," in the box.",(0,h.jsx)(o.ZP,{item:!0,xs:12,children:(0,h.jsx)(c.Z,{id:"retype-PVC",name:"retype-PVC",onChange:e=>{x(e.target.value)},label:"",value:C})})]})})}}}]);
//# sourceMappingURL=112.9ebf20ee.chunk.js.map