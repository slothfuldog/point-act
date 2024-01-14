import { Navigate } from "react-router-dom"

const Private = (props:{isLogin:any, element:JSX.Element})=>{
    console.log(props.isLogin)
    if(!props.isLogin){
        return <Navigate to='/login'/>
    }else{
        return props.element;
    }
}

export default Private