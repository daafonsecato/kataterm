import React from 'react'
import { useRoutes } from 'react-router-dom'
import StartView from '../Pages/Home/StartView'
import QuestionView from '../Pages/QuestionView/QuestionView'

export const AppRoutes = () => {
  const routes = useRoutes([
    { path: '/', element: <StartView /> },
    { path: '/quiz', element: <QuestionView /> }
  ])
  return routes
}