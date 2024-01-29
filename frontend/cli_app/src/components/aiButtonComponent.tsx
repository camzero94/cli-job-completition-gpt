const AiButtonComponent: React.FC = () => {
  return (
    <div className=' flex flex-col w-1/2 h-full item-center justify-center'>
      <div className=' flex  h-1/4 item-center justify-start mb-4'>
        <button
          type='button'
          data-te-ripple-init
          data-te-ripple-color='light'
          className='flex items-center rounded bg-gray-200 px-6 pb-2 pt-2.5 text-xs font-roboto uppercase leading-normal text-gray-600 shadow-[0_4px_9px_-4px_#3b71ca] transition duration-150 ease-in-out hover:bg-primary-600 hover:shadow-[0_8px_9px_-4px_rgba(59,113,202,0.3),0_4px_18px_0_rgba(59,113,202,0.2)] focus:bg-primary-600 focus:shadow-[0_8px_9px_-4px_rgba(59,113,202,0.3),0_4px_18px_0_rgba(59,113,202,0.2)] focus:outline-none focus:ring-0 active:bg-primary-700 active:shadow-[0_8px_9px_-4px_rgba(59,113,202,0.3),0_4px_18px_0_rgba(59,113,202,0.2)] dark:shadow-[0_4px_9px_-4px_rgba(59,113,202,0.5)] dark:hover:shadow-[0_8px_9px_-4px_rgba(59,113,202,0.2),0_4px_18px_0_rgba(59,113,202,0.1)] dark:focus:shadow-[0_8px_9px_-4px_rgba(59,113,202,0.2),0_4px_18px_0_rgba(59,113,202,0.1)] dark:active:shadow-[0_8px_9px_-4px_rgba(59,113,202,0.2),0_4px_18px_0_rgba(59,113,202,0.1)]'
        >
          <svg
            xmlns='http://www.w3.org/2000/svg'
            xmlns:xlink='http://www.w3.org/1999/xlink'
            version='1.1'
            id='Capa_1'
            x='0px'
            y='0px'
            viewBox='0 0 512 512'
            // style='enable-background:new 0 0 512 512;'
            xml:space='preserve'
            className='fill-current text-gray-500 w-5 h-5 mr-2'
          >
            <g>
              <polygon points='130.965,181.632 160.043,181.632 145.557,143.253  ' />
              <path d='M405.333,0H106.667C47.786,0.071,0.071,47.786,0,106.667v298.667C0.071,464.214,47.786,511.93,106.667,512h224.32   c3.477,0,6.912-0.277,10.347-0.512V405.333c0-35.346,28.654-64,64-64h106.155c0.235-3.435,0.512-6.869,0.512-10.347v-224.32   C511.93,47.786,464.214,0.071,405.333,0z M194.432,235.008c-6.892,2.579-14.572-0.903-17.173-7.787l-7.147-18.923h-49.301   l-7.232,18.965c-1.992,5.145-6.942,8.535-12.459,8.533c-1.618-0.009-3.222-0.305-4.736-0.875   c-6.884-2.615-10.344-10.316-7.729-17.199c0.002-0.006,0.004-0.011,0.006-0.017l42.133-110.592   c2.378-6.119,8.262-10.157,14.827-10.176l0,0c6.441,0.02,12.232,3.931,14.656,9.899l41.899,110.933   c2.624,6.893-0.837,14.608-7.73,17.232C194.441,235.004,194.437,235.006,194.432,235.008z M252.949,222.528   c0,7.364-5.97,13.333-13.333,13.333s-13.333-5.97-13.333-13.333V110.272c0-7.364,5.97-13.333,13.333-13.333   s13.333,5.97,13.333,13.333V222.528z' />
              <path d='M384,405.333v96.853c19.734-7.452,37.66-19.015,52.587-33.92l31.659-31.68c14.923-14.917,26.494-32.844,33.941-52.587   h-96.853C393.551,384,384,393.551,384,405.333z' />
            </g>
          </svg>
          Button
        </button>
      </div>
    </div>
  )
}

export default AiButtonComponent
