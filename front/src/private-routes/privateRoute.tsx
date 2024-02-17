import { Navigate } from "react-router-dom"

const Private = (props:{isLogin:any, element:JSX.Element})=>{
    console.log("PROS",props.isLogin)
    const currentPath = window.location.pathname
    console.log(currentPath)
    if(!props.isLogin && currentPath!== "/login"){
        console.log("masuk1")
        return <Navigate to='/login'/>
    } else if(!props.isLogin && currentPath == "/login"){
        console.log("masuk2")
        return props.element
    }
    else if(props.isLogin && currentPath != "/login"){
        console.log("masuk3")
        return props.element;
    }else{
        console.log("masuk4")
        return <Navigate to='/scoring' />
    }
}

export default Private