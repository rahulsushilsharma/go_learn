import { useEffect, useState } from 'react'
import './App.css'

function App() {
  const [users, setUser] = useState<{id:number, name:string}[]>([])
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
      console.log("Received:", event.data);
    };
  
    socket.onerror = (error) => {
      console.error("WebSocket error:", error);
    };
  
    socket.onclose = () => {
      console.log("WebSocket closed");
    };
  
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
      </div>
      
    </>
  )
}

export default App
