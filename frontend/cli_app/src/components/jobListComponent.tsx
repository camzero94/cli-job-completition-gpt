import JobComponent from './jobComponent'
import JobType from '../types/jobType'
import AiButtonComponent from './aiButtonComponent'

interface JobListProps {
  jobs: JobType.Job[]
}

const JobListComponent: React.FC<JobListProps> = ({jobs}) => {
  return jobs.map((job) => (
    <div className='flex w-full flex-row min-h-46 '>
      <JobComponent job={job}/>
      <AiButtonComponent/>
    </div>
  ))
}
export default JobListComponent
