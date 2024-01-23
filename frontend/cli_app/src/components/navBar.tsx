const NavBarComponent: React.FC = () => {
  return (
    <nav className=' bg-white border-gray-200 dark:bg-gray-900 '>
      <div className='flex  flex-row items-center justify-between w-full p-4 h-10vh '>
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
        <div className='flex justify-center w-1/4' id='logo-name-container-el'>
          <ul className='font-medium  flex flex-col items-center p-4 md:p-0 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 rtl:space-x-reverse md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700'>
            <li>
              <img
                className='object-fit h-12 w-70'
                src='https://www.freepnglogos.com/uploads/spotify-logo-png/spotify-download-logo-30.png'
                alt='spotify logo'
              />
            </li>
            <li>
              <h3>AiResume</h3>
            </li>
          </ul>
        </div>

        <div className='w-1/4' id='log-sign-container-el'>
          <ul className='font-medium flex  items-center justify-center p-4 md:p-4 sm:p-4 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 rtl:space-x-reverse md:mt-0 md:border-0 md:bg-white dark:bg-gray-800 md:d'>
            <li>Sign In</li>
            <li>Sign Up</li>
          </ul>
        </div>
      </div>
    </nav>
  )
}

export default NavBarComponent
