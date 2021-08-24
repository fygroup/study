### GOJS
```
GoJS是一个实现交互类图表（比如流程图，树图，关系图，力导图等等）的JS库

结构

diagram

//视图层
Node  Link

//数据层
modelArray  linkArray


//用法
<div id="myDiagramDiv" style="width:400px; height:150px; background-color: #DAE4E4;"></div>

import GOJS from go.js

var GOJS = go.GraphObject.make
var myDiagram = GOJS(go.Diagram, "myDiagramDiv", {
    //图的一些特征
    //key: value
})

//渲染层(NodeTemplate、LinkTemplate)
myDiagram.nodeTemplate =
  $(go.Node, "Horizontal",
    { background: "#44CCFF" },
    $(go.Picture,
      { margin: 10, width: 50, height: 50, background: "red" },
      new go.Binding("source")),
    $(go.TextBlock, "Default Text",
      { margin: 12, stroke: "white", font: "bold 16px sans-serif" },
      new go.Binding("text", "name"))
  );
myDiagram.linkTemplate =
  $(go.Link,
    { routing: go.Link.Orthogonal, corner: 5 },
    $(go.Shape, { strokeWidth: 3, stroke: "#555" })); // the link shape


//数据层(Model、TreeModel、GraphLinksModel、....)
var model = $(go.Model);
var model = $(go.GraphLinksModel);
var model = $(go.TreeModel);
model.nodeDataArray =
[
  { key: "1",              name: "Don Meow",   source: "cat1.png" },
  { key: "2", parent: "1", name: "Demeter",    source: "cat2.png" },
  { key: "3", parent: "1", name: "Copricat",   source: "cat3.png" },
  { key: "4", parent: "3", name: "Jellylorum", source: "cat4.png" },
  { key: "5", parent: "3", name: "Alonzo",     source: "cat5.png" },
  { key: "6", parent: "2", name: "Munkustrap", source: "cat6.png" }
];
model.nodeDataArray =
[
  { key: "A" },
  { key: "B" },
  { key: "C" }
];
model.linkDataArray =
[
  { from: "A", to: "B" },
  { from: "B", to: "C" }
];
myDiagram.model = model;
或者
var nodeDataArray = [...]
var linkDataArray = [...]
diagram.model = new go.GraphLinksModel(nodeDataArray, linkDataArray);

```

