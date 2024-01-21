import { useState } from "react";

const Login = () =>{
    const [pass, setPass] = useState(false);
    const handleCheck = ()=>{
        setPass(!pass);
    }
    return(
        <div className="flex flex-row justify-center items-center">
            <div className=" border-solid border-2 border-white shadow-xl w-2/6 mt-20 bg-slate-50">
            <p className="flex justify-center mt-5 text-xl font-semibold">Login</p>
            <hr className="ml-5 mr-5 mt-5"/>
            <div className="ml-5 mt-5 mb-5 mr-5">
            <p>Username</p>
            <input type="text" className="rounded-sm mt-2 w-full "/>
            <p>Password</p>
            <input type={!pass ? "password" : "text"} className="rounded-sm mt-2 w-full"/><br/>
            <div className="flex flex-row mt-2 ">
            <input id="check_pass" className="hover:cursor-pointer" type="checkbox" onChange={() => {handleCheck()}} /> <label htmlFor="check_pass" className="ml-2 text-sm hover:cursor-pointer">Show password</label>
            </div>
            <button className="p-1 rounded-md w-full mt-5 btn-def" >Login</button>
            </div>
            </div>
        </div>
    )
}

export default Login;