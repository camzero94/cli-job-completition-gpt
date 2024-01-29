import JobListComponent from '../components/jobListComponent'
import JobType from '../types/jobType'
import { useState, useEffect, useContext} from 'react'
import { Job_ctx,IContext } from '../store/context/contextApp'
import SearchComponent from '../components/searchComponent'


function Home() {

  const {jobs, setJobs,jobKeyword,skills,page,sendGetReq, setSendGetReq} = useContext(Job_ctx) as IContext

  const fetchJobsAsync = async (job: JobType.JobQuery) => {
    // Create Request Get Object
    const { jobQuery, skillsQuery, page } = job
    const getRequObj = new Request(
      `http://localhost:3000/getJobs?myJob=${jobQuery}&skills=${skillsQuery}skill1&pages=${page}`,
      {
        method: 'GET',
        headers: new Headers({
          'Content-Type': 'application/json',
          'Access-Control-Allow-Origin': '*'
        }),
      }
    )
    try {
      const res = await fetch(getRequObj)
      const data = await res.json()
      console.log(data)
      // setJobs(data)
      setSendGetReq(false)
    } catch (err) {}
  }

  useEffect(() => {
    const jobQuery:JobType.JobQuery = {
      jobQuery: jobKeyword,
      skillsQuery: skills,
      page: page,
    }
    console.log(jobQuery)
    sendGetReq ? fetchJobsAsync(jobQuery) : null

  }, [sendGetReq])

  return (
    <>
      <SearchComponent />
      <JobListComponent jobs={[]} />
    </>
  )
}

export default Home
