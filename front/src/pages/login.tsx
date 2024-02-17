import { useEffect, useState } from "react";
import Axios from "axios";
import { useNavigate } from "react-router-dom";

const Login = (props:{isLogin:any}) =>{
    const navigate = useNavigate()
    const [pass, setPass] = useState(false);
    const [user, setUser] = useState("");
    const [pwd, setPwd] = useState("");
    const [login, setLogin] = useState([]);
    const handleCheck = ()=>{
        setPass(!pass);
    }

    useEffect ( () =>{
        if(props.isLogin){
            navigate("/scoring")
            window.location.reload()
        }
    }, [])
    const Login = () =>{
            Axios.post("http://localhost:3000/login",{
                username: user,
                password: pwd
            })
            .then((res) =>{
                if(res.status === 200){
                    localStorage.setItem("login", JSON.stringify(res.data.response))
                    navigate("/scoring")  
                    window.location.reload()
                }
            }
            )
        .catch((e) =>{
            console.log(e)
            alert("DATA NOT FOUND")
        })
    }

    return(
        <div className="flex flex-row justify-center items-center">
            <form className=" border-solid border-2 border-white shadow-xl w-2/6 mt-20 bg-slate-50" onSubmit={e => e.preventDefault()}>
            <p className="flex justify-center mt-5 text-xl font-semibold">Login</p>
            <hr className="ml-5 mr-5 mt-5"/>
            <div className="ml-5 mt-5 mb-5 mr-5">
            <p>Username</p>
            <input type="text" className="rounded-sm mt-2 w-full " onChange={(e) => setUser(e.target.value)}/>
            <p>Password</p>
            <input type={!pass ? "password" : "text"} className="rounded-sm mt-2 w-full" onChange={(e) => setPwd(e.target.value)}/><br/>
            <div className="flex flex-row mt-2 ">
            <input id="check_pass" className="hover:cursor-pointer" type="checkbox" onChange={() => {handleCheck()}} /> <label htmlFor="check_pass" className="ml-2 text-sm hover:cursor-pointer">Show password</label>
            </div>
            <button type="submit" className="p-1 rounded-md w-full mt-5 btn-def" onClick={() => Login()}>Login</button>
            </div>
            </form>
        </div>
    )
}

export default Login;