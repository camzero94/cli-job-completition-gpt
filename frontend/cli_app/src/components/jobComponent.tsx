import JobType from '../types/jobType'
interface JobComponentProps {
  job: JobType.Job
}

const JobComponent: React.FC<JobComponentProps> = ({ job }) => {
  const { jobName, company, content, link } = job || {}
  return (
    <div className='flex w-full item-center justify-center mt-8'>
      <div className=' sm:w-4/5 md:w-3/4 xl:w-1/2 min-w-72  h-5/6    rounded-full border border-gray-400  bg-white p-6 flex flex-col  shadow-lg transition duration-300 ease-in-out hover:scale-110'>
        <p className='font-roboto  font-weight-500 text-sm text-gray-600 flex items-center'>
          <svg
            className='fill-current text-gray-500 w-3 h-3 mr-2'
            xmlns='http://www.w3.org/2000/svg'
            viewBox='0 0 24 24'
          >
            <path d='M19,4h-1.1c-.46-2.28-2.48-4-4.9-4h-2c-2.41,0-4.43,1.72-4.9,4h-1.1C2.24,4,0,6.24,0,9v10c0,2.76,2.24,5,5,5h14c2.76,0,5-2.24,5-5V9c0-2.76-2.24-5-5-5ZM11,2h2c1.3,0,2.4,.84,2.82,2h-7.63c.41-1.16,1.51-2,2.82-2Z' />
          </svg>
          {jobName}
        </p>

        <div className='font-lato text-gray-900 text-sm mb-2'>{content}</div>

        <div class='px-6 pb-2 font-roboto text-xs flex justify-center'>
          <span class='inline-block bg-gray-200 rounded-full px-3 py-1 font-semibold text-gray-700 mr-2 mb-2'>
            #photography
          </span>
          <span class='inline-block bg-gray-200 rounded-full px-3 py-1 font-semibold text-gray-700 mr-2 mb-2'>
            #travel
          </span>
          <span class='inline-block bg-gray-200 rounded-full px-3 py-1 font-semibold text-gray-700 mr-2 mb-2'>
            #winter
          </span>
        </div>
      </div>
    </div>
  )
}

export default JobComponent
