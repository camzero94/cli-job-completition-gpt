import LayoutComponent from './components/layoutComponent'
import NavBarComponent from './components/navBar'
import Home from './pages/Home'
import JobContextProvider from './store/context/contextApp' 

function App() {
  return (
    <JobContextProvider>
      <LayoutComponent>
        <NavBarComponent />
        <Home />
      </LayoutComponent>
    </JobContextProvider>
  )
}

export default App
