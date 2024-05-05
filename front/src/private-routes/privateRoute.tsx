import { Navigate } from "react-router-dom"

const Private = (props:{isLogin:any, element:JSX.Element})=>{
    const currentPath = window.location.pathname
    if(!props.isLogin && currentPath!== "/login"){
        return <Navigate to='/login'/>
    } 
    else if(currentPath === "/"){
        return <Navigate to ='/login' />
    }
    else if(!props.isLogin && currentPath == "/login"){
        return props.element
    }
    else if(props.isLogin && currentPath != "/login"){
        return props.element;
    }else{
        return <Navigate to='/scoring' />
    }
}

export default Private