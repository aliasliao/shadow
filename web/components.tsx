import * as React from 'react'
import { loadingContainer } from './styles'

export function Loading() {
  return (
    <div className={loadingContainer}>
      <div></div>
      <div></div>
      <div></div>
      <div></div>
    </div>
  )
}
