import {
    FormControl,
    FormLabel,
    Select,
    Skeleton
} from "@chakra-ui/react"
import useRoles from "../../hooks/api/useRoles"

export default function RolSelect(props) {
    const { onChange, value } = props

    const { query } = useRoles()
    const { isPending, data } = query


    if (isPending) {
        return (
            <FormControl isRequired {...props}>
                <FormLabel>User Rol</FormLabel>
                <Skeleton height='20px' fadeDuration={1} mt={5} />
            </FormControl>
        )
    }


    return (
        <FormControl isRequired {...props}>
            <FormLabel>User Rol</FormLabel>
            <Select size="sm" placeholder="Select Rol" variant="filled" onChange={onChange} value={value}>
                {data?.data?.roles.map((item) => {
                    return <option key={item} value={item}>{item}</option>
                })}
            </Select>
        </FormControl>
    )
}