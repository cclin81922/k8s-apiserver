

-----------
# Lin v1 cclin

>bdocs-tab:example Lin Example

```bdocs-tab:example_yaml

apiVersion: cclin.cclin/v1
kind: Lin
metadata:
  name: lin-example
spec:


```


Group        | Version     | Kind
------------ | ---------- | -----------
cclin | v1 | Lin







Lin

<aside class="notice">
Appears In:

<ul> 
<li><a href="#linlist-v1-cclin">LinList cclin/v1</a></li>
</ul> </aside>

Field        | Description
------------ | -----------
apiVersion <br /> *string*    | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
kind <br /> *string*    | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
metadata <br /> *[ObjectMeta](#objectmeta-v1-meta)*    | 
spec <br /> *[LinSpec](#linspec-v1-cclin)*    | 
status <br /> *[LinStatus](#linstatus-v1-cclin)*    | 


### LinSpec v1 cclin

<aside class="notice">
Appears In:

<ul>
<li><a href="#lin-v1-cclin">Lin cclin/v1</a></li>
</ul></aside>

Field        | Description
------------ | -----------

### LinStatus v1 cclin

<aside class="notice">
Appears In:

<ul>
<li><a href="#lin-v1-cclin">Lin cclin/v1</a></li>
</ul></aside>

Field        | Description
------------ | -----------

### LinList v1 cclin



Field        | Description
------------ | -----------
apiVersion <br /> *string*    | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
items <br /> *[Lin](#lin-v1-cclin) array*    | 
kind <br /> *string*    | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
metadata <br /> *[ListMeta](#listmeta-v1-meta)*    | 





