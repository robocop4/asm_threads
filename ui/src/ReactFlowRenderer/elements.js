import ReactFlow, {
  useNodesState,
  useEdgesState,
  addEdge,
  MiniMap,
  Controls,
  Background,
  Node,
  Edge,
  Position,
  ConnectionMode,
  MarkerType,
} from 'reactflow';
export const nodes = [
    {
      id: "1",
    //  type: "input",
      data: {
        label: "Node 1",
      },
      position: { x: 250, y: 0 },
      type: 'bidirectional',
      // sourcePosition: Position.Left,
      // targetPosition: Position.Right,
    },
    {
      id: "4",
     // type: "input",
      data: {
        label: "Node 4",
      },
      position: { x: 500, y: 0 },
      type: 'bidirectional',
      // sourcePosition: Position.Right,
      // targetPosition: Position.Left,
    },
    {
      id: "2",
      data: {
        label: "Node 2",
      },
      position: { x: 100, y: 100 },
      type: 'bidirectional',
      // sourcePosition: Position.Right,
      // targetPosition: Position.Left,
    },
    {
      id: "3",
      data: {
        label: "Node 3",
      },
      position: { x: 400, y: 100 },
      style: {
        background: "#D6D5E6",
        color: "#333",
        border: "1px solid #222138",
        width: 180,
      },
      type: 'bidirectional',
      // sourcePosition: Position.Right,
      // targetPosition: Position.Left,
      // sourcePosition: Position.Left,
      // targetPosition: Position.Right,
    },
  ];
  
  export const edges = [
    { id: "e1-2", step: "0", source: "1", target: "2", type: "bidirectional", animated: false,  sourceHandle: 'left', targetHandle: 'right', markerEnd: { type: MarkerType.Arrow },  style: {strokeWidth: 2, stroke: 'rgba(0, 0, 255, 0.5)',}},
    { id: "e1-3", step: "1", source: "2", target: "3", type: "bidirectional", animated: false,  sourceHandle: 'left', targetHandle: 'right', markerEnd: { type: MarkerType.Arrow },  style: {strokeWidth: 2, stroke: 'rgba(0, 0, 255, 0.5)',}},
    { id: "e1-5", step: "2", source: "3", target: "1", type: "bidirectional", animated: false,  sourceHandle: 'left', targetHandle: 'right', markerEnd: { type: MarkerType.Arrow },  style: {strokeWidth: 2,  stroke: 'rgba(0, 0, 255, 0.5)',}},
    { id: "e1-4", step: "3", source: "1", target: "4", type: "bidirectional", animated: false,  sourceHandle: 'left', targetHandle: 'right', markerEnd: { type: MarkerType.Arrow },  style: {strokeWidth: 2, stroke: 'rgba(0, 0, 255, 0.5)',}},
  
    

    { id: "e2-2", step: "0", source: "1", target: "2", type: "bidirectional", animated: false,  sourceHandle: 'left', targetHandle: 'right', markerEnd: { type: MarkerType.Arrow },  style: {strokeWidth: 2, stroke: 'rgba(255, 255, 0, 0.5)',}},
    { id: "e2-3", step: "1", source: "2", target: "3", type: "bidirectional", animated: false,  sourceHandle: 'left', targetHandle: 'right', markerEnd: { type: MarkerType.Arrow },  style: {strokeWidth: 2, stroke: 'rgba(255, 255, 0, 0.5)'}},
    { id: "e2-5", step: "2", source: "3", target: "1", type: "bidirectional", animated: false,  sourceHandle: 'left', targetHandle: 'right', markerEnd: { type: MarkerType.Arrow },  style: {strokeWidth: 2, stroke: 'rgba(255, 255, 0, 0.5)'}},
    { id: "e2-4", step: "3", source: "1", target: "3", type: "bidirectional", animated: false,  sourceHandle: 'left', targetHandle: 'right', markerEnd: { type: MarkerType.Arrow },  style: {strokeWidth: 2, stroke: 'rgba(255, 255, 0, 0.5)'}},
  
  ];


  
  