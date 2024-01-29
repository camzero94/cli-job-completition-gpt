const NavBarComponent: React.FC = () => {
  return (
    <nav className=' w-full bg-white border-gray-200 dark:bg-gray-900 py-10'>
      <div className='flex flex-row items-center justify-between mb-10 h-10vh'>
        <a
          href='https://flowbite.com/'
          className='flex items-center space-x-3 rtl:space-x-reverse'
        >
          <img
            className='object-fit h-12 w-70'
            src='https://www.freepnglogos.com/uploads/spotify-logo-png/spotify-download-logo-30.png'
            alt='spotify logo'
          />
          <span>AiResume</span>
        </a>
          <ul className='font-medium  flex space-x-3 items-center p-4  md:space-x-8 md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700'>
              <img
                className='object-fit h-12 w-70'
                src='https://www.freepnglogos.com/uploads/spotify-logo-png/spotify-download-logo-30.png'
                alt='spotify logo'
              />
              <h3>AiResume</h3>
          </ul>
          <ul className='font-medium flex  items-center  sm:p-4  space-x-2 rtl:space-x-reverse md:mt-0 md:border-0 md:bg-white dark:bg-gray-800'>
            <li>Sign In</li>
            <li>Sign Up</li>
          </ul>
      </div>
    </nav>
  )
}

export default NavBarComponent
