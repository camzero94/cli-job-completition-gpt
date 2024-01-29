import { useRef,useContext } from 'react'
import { IContext,Job_ctx} from '../store/context/contextApp'


const SearchKeywordComponent: React.FC = () => {

  const jobKeywordRef = useRef<HTMLInputElement>(null)
  const {setJobKeyword, setSendGetReq,setSkills,skills} = useContext(Job_ctx) as IContext

  // Set the keyword value to the context
  const handleKeywordSubmit = () => {
    if (jobKeywordRef.current ) {
      console.log(jobKeywordRef.current.value)
      console.log(skills)
      setJobKeyword(jobKeywordRef.current.value)
      // setSkills(prev=>[...prev,...skills])
      setSendGetReq(true)
    }
  }
  return (
    <div className='flex w-full justify-center'>
      <form className='w-1/2'>
        <label
          for='default-search'
          className='mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white'
        >
          Search
        </label>
        <div className='relative'>
          <div className='absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none'>
            <svg
              className='fill-current text-gray-500 w-3 h-3 mr-2'
              xmlns='http://www.w3.org/2000/svg'
              viewBox='0 0 24 24'
            >
              <path d='M19,4h-1.1c-.46-2.28-2.48-4-4.9-4h-2c-2.41,0-4.43,1.72-4.9,4h-1.1C2.24,4,0,6.24,0,9v10c0,2.76,2.24,5,5,5h14c2.76,0,5-2.24,5-5V9c0-2.76-2.24-5-5-5ZM11,2h2c1.3,0,2.4,.84,2.82,2h-7.63c.41-1.16,1.51-2,2.82-2Z' />
            </svg>
          </div>
          <input
            type='search'
            ref={jobKeywordRef}
            id='default-search'
            className='block w-full p-4 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
            placeholder='Search Job Keyword...'
            required
          />
          <button
            type='button'
            className='text-white absolute end-2.5 bottom-2.5 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800'
            onClick={handleKeywordSubmit} 
          >
            Search
          </button>
        </div>
      </form>
    </div>
  )
}

export default SearchKeywordComponent
