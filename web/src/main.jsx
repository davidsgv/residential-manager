import React from 'react'
import ReactDOM from 'react-dom/client'
import { ChakraProvider } from '@chakra-ui/react'
import App from './App.jsx'
import "./assets/css/index.css"

const toastOptions = {
  defaultOptions: {
    position: 'top-right',
    duration: 10000,
    isClosable: true
  }
}

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <ChakraProvider toastOptions={toastOptions}>
      <App />
    </ChakraProvider>
  </React.StrictMode>
)