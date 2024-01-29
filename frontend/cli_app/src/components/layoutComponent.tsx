
interface LayoutComponentProps {
  children: React.ReactNode
}
const LayoutComponent: React.FC<LayoutComponentProps> = ({children}) => {
  return (
      <div className='flex flex-col w-full h-screen px-12  items-center '>
      {children}
      </div>
  )
}

export default LayoutComponent 
