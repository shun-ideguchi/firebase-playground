import type { FC } from 'react'
import { Home } from './Home'

const App: FC = () => {
  return (
    <>
      <div className="App">
        <h1>Firebase Playground</h1>
        <Home />
      </div>
    </>
  )
}

export default App
