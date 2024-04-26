import { useState } from 'react'
import validate from '../helpers/validations';

export default function useInput(initialValue, validationType, disabled){
    const [value, setValue] = useState(initialValue)
    const [error, setError] = useState(false)

    const handleChange = (event) => {
        const value = event.target.value 
        const error = !validate(value, validationType)

        setValue(value);
        if (value == ""){
            setError(false)
            return
        }
        setError(error)
    };

    return {setValue, value, error, handleChange, empty: value.length == 0, disabled}
}