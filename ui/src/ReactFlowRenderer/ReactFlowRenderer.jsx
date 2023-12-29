import React, { useState, useCallback, useEffect, useRef } from "react";
import ReactFlow, {
  ReactFlowProvider,
  addEdge,
  useNodesState,
  useEdgesState,
  ConnectionLineType,
  Controls,
  Background,
  MiniMap,
} from "react-flow-renderer";

import Grid from '@mui/material/Grid';



import { nodes as initialNodes, edges as initialEdges } from "./elements";
import { Button, Modal, Input, Form } from "antd";





function ReactFlowRenderer(props) {



 
//  const [nodes, setNodes, onNodesChange] = useNodesState(initialNodes);
 // const [edges, setEdges, onEdgesChange] = useEdgesState(initialEdges);
 const reactFlowWrapper = useRef(null);
  const [nodes, setNodes, onNodesChange] = useNodesState(initialNodes);
  const [edges, setEdges, onEdgesChange] = useEdgesState(initialEdges); 
  const [currentEdgeIndex, setCurrentEdgeIndex] = useState(0);
  const [position, setPosition] = useState({ x: 0, y: 0 });

  const [file1, setFile1] = useState(null);
  const [file2, setFile2] = useState(null);
  const [maxSteps, setMaxSteps] = useState(0);
  const [file3, setFile3] = useState(null);
  const [file1Content, setFile1Content] = useState('');
  const [file2Content, setFile2Content] = useState([]);

  const desiredX = 100; // Replace with your desired x coordinate
  const desiredY = 100; // Replace with your desired y coordinate

  useEffect(() => {

    console.log(position);
    setPosition({ x: desiredX, y: desiredY });
  }, [desiredX, desiredY]);

//   useEffect(() => {
//   if (props.file1.length !== 0 && props.file2.length !==0 ) {

//     const base = JSON.parse(props.file1);
//     const trace = JSON.parse(props.file2);
//     setNodes(base);
//     setEdges(trace);
//   } 
// }, []);


const exportButtonStyle = {
  position: 'absolute',
  right: '12px',
  top: '12px',
  minWidth: '120px',
  backgroundColor: '#F5F5F5',
  zIndex: 5,
  border: '1px solid #CACACA',
  padding: '12px',
  borderRadius: '5px',
};





const handleFile1Change = (event) => {
  const selectedFile = event.target.files[0];
  setFile1(selectedFile);
};

const handleFile2Change = (event) => {
  const selectedFile = event.target.files[0];
  setFile2(selectedFile);
};



const maxStep = edges.reduce((max, edge) => {
  const step = parseInt(edge.step); // Преобразуем значение "step" в число
  return step > max ? step : max;
}, -Infinity); // Начальное значение - отрицательная бесконечность


//Нажата кнопка upload
const handleReadFiles = () => {
  if (file1) {
    const reader1 = new FileReader();
    reader1.onload = (event) => {

      //TODO:
      const jsonObject = JSON.parse(event.target.result);
      setFile1Content(jsonObject);
      setNodes(jsonObject);
    };
    reader1.readAsText(file1);
  }

  if (file2) {
    const reader2 = new FileReader();
    reader2.onload = (event) => {
      //TODO:
      const jsonObject = JSON.parse(event.target.result);
      setFile2Content(jsonObject);
      
     setEdges(jsonObject);

    //смотрим самый большой step
    const maxStep = jsonObject.reduce((max, edge) => {
      const step = parseInt(edge.step); // Преобразуем значение "step" в число
      return step > max ? step : max;
    }, -Infinity); // Начальное значение - отрицательная бесконечность
    


    //обновляем количество шагов в программе
    setMaxSteps(maxStep);
    //обнуляем текущий шаг
   
    //console.log('Самый большой step:', );

    };
    reader2.readAsText(file2);
  }




};








  //const getNodeId = () => Math.random();
  function onInit() {
   // console.log("Logged");
  }




    //+1
   function displayCustomNamedNodeModal() {
     // Create a copy of the current edges array
     const updatedEdges = [...edges];
     // Update the animated property of the edge at currentEdgeIndex
     updatedEdges.forEach((edge) => {
      if (edge.step == currentEdgeIndex) {
        edge.animated = true;
      } else {
        edge.animated = false;
      }
    });
    
    // updatedEdges[currentEdgeIndex].animated = true;
     // Set the updated edges array using the setEdges function
     setEdges(updatedEdges);
     console.log(updatedEdges)
     console.log(currentEdgeIndex);
     //scrollToCoordinates(0, 0);
     setPosition({ x: 0, y: 0 });
     setCurrentEdgeIndex((prevIndex) => (prevIndex + 1) % maxStep)
  }
  // function handleCancel() {
  //   setIsModalVisible(false);
  // }
  function handleOk(data) {
 
  }


  // const onAddEdges = useCallback(
  //   (data) => {
  //     const newNode = {
  //       id: String(getNodeId()),
  //       data: { label: data },
  //       position: {
  //         x: 50,
  //         y: 0,
  //       },
  //     };
  //     setNodes((nds) => nds.concat(newNode));
  //   },
  //   [setNodes]
  // );




  return (
    <div style={{ height: "100vh"}}>
      {/* <Modal
        title="Basic Modal"
       // visible={isModalVisible}
       // onCancel={handleCancel}
      >



        <Form onFinish={handleOk} autoComplete="off" name="new node">
          <Form.Item label="Node Name" name="nodeName">
            <Input />
          </Form.Item>

          <Form.Item>

         
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit">
              Submit
            </Button>
          </Form.Item>
        </Form>
      </Modal> */}
     


   

<ReactFlowProvider>


{/* <div  className={styles.workflowWrapper}> */}
        <div  style={exportButtonStyle}>


        <Grid container spacing={2}>
      <Grid item xs={12}>
        <>
          <input  type="file" onChange={handleFile1Change} /> 
          <input  type="file" onChange={handleFile2Change} /> 
          <button onClick={handleReadFiles} >Upload</button>
        </>
      </Grid>
      <Grid item xs={12}>
        <>
            <button onClick={displayCustomNamedNodeModal} >Next</button>
        </>
      </Grid>
    
    </Grid>



          </div>
      <ReactFlow
        nodes={nodes}
        edges={edges}
        onNodesChange={onNodesChange}
        onEdgesChange={onEdgesChange}
        //onConnect={onConnect}
        onInit={onInit}
        fitView
        nodeOrigin={[1,1]}
        attributionPosition="bottom-left"
        connectionLineType={ConnectionLineType.SmoothStep}
        
      />

<Background />

{/* <MiniMap /> */}
      <Controls />
     



      




   {/* </div> */}

      </ReactFlowProvider>


    </div>
  );
}

export default ReactFlowRenderer;
