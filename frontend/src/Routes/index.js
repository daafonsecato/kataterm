import React from 'react'
import { useRoutes } from 'react-router-dom'
import StartView from '../Pages/Home/StartView'
import QuestionView from '../Pages/QuestionView/QuestionView'
import ResultView from '../Pages/ResultView/ResultView'

export const AppRoutes = () => {
  const routes = useRoutes([
    { path: '/', element: <StartView /> },
    { path: '/quiz', element: <QuestionView /> },
    { path: '/result', element: <ResultView /> }
  ])
  return routes
}