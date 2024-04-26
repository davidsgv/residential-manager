import { useState } from "react";
import { useNavigate } from '@tanstack/react-router'
import { getClaims, setToken } from "../helpers/jwt";

export default function useToken(){
    const [claims, setClaims] = useState(getClaims())
    const navigate = useNavigate()

    const updateToken = (token) => {
        setToken(token)
        setClaims(getClaims())
        navigate({ to: '/app/'})
    }

    return [claims, updateToken]
}