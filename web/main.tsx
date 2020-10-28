import * as React from 'react'
import * as ReactDOM from 'react-dom'
import App from "./app";
import { initializeNotification } from './utils'

initializeNotification()

ReactDOM.render(
  <App/>,
  document.getElementById('shadow-app'),
)
