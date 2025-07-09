import { useEffect, useState } from 'react'
import './App.css'

function App() {
  const [users, setUser] = useState<any[]>([])
  async function getUsers(){
    const user = await fetch('http://localhost:8000/users')
    const json = await user.json()
    console.log(json)
    setUser(json)
  }
  useEffect(()=>{
    getUsers()
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
