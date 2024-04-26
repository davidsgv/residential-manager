import {
    Tr,
    Td,
    Skeleton
} from '@chakra-ui/react'

export default function TableSkeleton() {
    return (
        <>
            <Skeleton height='20px' fadeDuration={1} mt={5} />
            <Skeleton height='20px' fadeDuration={1} mt={5} />
            <Skeleton height='20px' fadeDuration={1} mt={5} />
            <Skeleton height='20px' fadeDuration={1} mt={5} />
        </>
    )
}