import { useEffect, useState } from 'react'
import './App.css'

function App() {
  const [users, setUser] = useState<{id:number, name:string}[]>([])
  const [socketConn, setSocketConn] = useState<WebSocket>()
  const [gameData, setGameData] = useState("{}")
  async function getUsers(){
    const user = await fetch('http://localhost:8000/users')
    const json = await user.json()
    console.log(json)
    setUser(json)
  }
  useEffect(() => {
    const socket = new WebSocket("ws://localhost:8000/ws");
  
    
    socket.onopen = () => {
      console.log("Connected to WebSocket");
    };
  
    socket.onmessage = (event) => {
      setGameData(event.data)
    };
  
    socket.onerror = (error) => {
      console.error("WebSocket error:", error);
    };
  
    socket.onclose = () => {
      console.log("WebSocket closed");
    };
    setSocketConn(socket)
  
    return () => {
      socket.close();
    };
  }, []);
  
  // function socketConnection(){
  //   const exampleSocket = new WebSocket(
  //     "http://localhost:8000/ws",
  //     "protocolOne",
  //   );
  //   console.log(exampleSocket)
  //   exampleSocket.onopen = () => {
  //     console.log("Connected");
  //   };
    
  //   exampleSocket.onmessage = (event) => {
  //     console.log("Message from server:", event.data);
  //   };
    
  //   exampleSocket.onerror = (err) => {
  //     console.error("WebSocket error:", err);
  //   };
    
  //   exampleSocket.onclose = () => {
  //     console.log("WebSocket closed");
  //   };
  // }


  useEffect(()=>{
    getUsers()
    // socketConnection()
  },[])
  return (
    <>
      <div>
        {
          users.map(ele=>{
            return <div key={ele?.id}>
            <p>{ele.name}</p>
            </div>
          })


        }

        <button onClick={()=>{
          socketConn?.send('data')
        }}>add</button>

<GameBoard gameData={gameData}/>




      </div>
      
    </>
  )
}


function GameBoard(props:{gameData : string}){
  const jsonData = JSON.parse(props.gameData)
  const snakePos = jsonData?.SnakePosition?.map((val: number[])=>val[0]+''+val[1])

  function color(i:number,j:number){
    let background = "green"
    if(jsonData?.Food_x == i && jsonData?.Food_y == j){
      background = ""
    }
    if (snakePos.includes(j+""+i)){
      background = "blue"
    }
    return background
  }

  return <>
  <p>{props.gameData}</p>


  <div>
    
    {
      Array.from({length:jsonData?.Bound_y}).map((_,i)=><>
      <div key={i} style={{ height:"5px", width:'100%' ,display:"flex"}}>
        {
          Array.from({length:jsonData?.Bound_x}).map((_,j)=>{
          const background = color(j,i)
          return <>
          <div key={j}  style={{ width:"5px", background:background}}>
            
          </div>
          </>})

        }
      </div>
      </>)
    }
  </div>
  </>
}

export default App
