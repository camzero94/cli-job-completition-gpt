import { useRef, useContext } from 'react'
import { IContext, Job_ctx } from '../store/context/contextApp'

const SkillComponent: React.FC = () => {
  const skillsRef = useRef<HTMLInputElement>(null)
  const { setSkills } = useContext(Job_ctx) as IContext

  const handleAddSkill = () => {
    if (skillsRef.current) {
      let skills = skillsRef.current.value.split(' ')
      setSkills(skills)
      console.log('add skill')
    }
  }

  return (
    <div className='flex w-full justify-center'>
      <form className='w-2/5'>
        <label
          for='default-search'
          className='mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white'
        >
          Search
        </label>
        <div className='relative'>
          <div className='absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none'>
            <svg
              xmlns='http://www.w3.org/2000/svg'
              className='fill-current text-gray-500 w-4 h-4 mr-2'
              id='Layer_1'
              data-name='Layer 1'
              viewBox='0 0 24 24'
            >
              <path d='m14,23c0,.552-.448,1-1,1H1c-.552,0-1-.448-1-1,0-3.866,3.134-7,7-7s7,3.134,7,7ZM7,6c-2.209,0-4,1.791-4,4s1.791,4,4,4,4-1.791,4-4-1.791-4-4-4Zm17-1v8c0,2.761-2.239,5-5,5h-4.526c-.945-1.406-2.275-2.533-3.839-3.227,1.437-1.096,2.365-2.826,2.365-4.773,0-3.314-2.686-6-6-6-1.084,0-2.102.288-2.979.791.112-2.658,2.294-4.791,4.979-4.791h10c2.761,0,5,2.239,5,5Zm-4,10c0-.553-.448-1-1-1h-3.5c-.552,0-1,.447-1,1s.448,1,1,1h3.5c.552,0,1-.447,1-1Z' />
            </svg>
          </div>
          <input
            type='search'
            id='default-search'
            className='block w-full p-4 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
            placeholder='Select Skills...'
            ref={skillsRef}
            required
          />
          <button
            type='button'
            className='text-white absolute end-2.5 bottom-2.5 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800'
            onClick={handleAddSkill}
          >
            Add
          </button>
        </div>
      </form>
    </div>
  )
}
export default SkillComponent
