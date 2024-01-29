import React, { createContext, useState } from 'react'
import { Dispatch, SetStateAction } from 'react'
import JobType from '../../types/jobType'

export interface IContext {
  jobs: JobType.Job[]
  setJobs: Dispatch<SetStateAction<JobType.Job[]>>
  jobKeyword: string
  setJobKeyword: Dispatch<SetStateAction<string>>
  skills: string[]
  setSkills: Dispatch<SetStateAction<string[]>>
  page: number 
  setPage: Dispatch<SetStateAction<number >>

  //Manage submits
  sendGetReq: boolean
  setSendGetReq: Dispatch<SetStateAction<boolean>>

  //Add Skills States
  addSkill: boolean
  setAddSkill: Dispatch<SetStateAction<boolean>>
}

export const Job_ctx= createContext<IContext | null>(null)

interface IProps {
  children: React.ReactNode
}

const JobContextProvider: React.FC <IProps> =  ({ children }) => {

  const [jobs, setJobs] = useState<JobType.Job[]>([])
  const [jobKeyword, setJobKeyword] = useState<string>('')
  const [skills, setSkills] = useState<string[]>([])
  const [page, setPage] = useState<number >(1)
  const [sendGetReq, setSendGetReq] = useState<boolean>(false)
  const [addSkill, setAddSkill] = useState<boolean>(false)


  const statesPage = {
    jobs: jobs,
    setJobs: setJobs,
    jobKeyword: jobKeyword,
    setJobKeyword: setJobKeyword,
    skills: skills,
    setSkills: setSkills,
    page: page,
    setPage: setPage,
    sendGetReq: sendGetReq,
    setSendGetReq: setSendGetReq,
    addSkill: addSkill,
    setAddSkill: setAddSkill,
  }

  return (
    <Job_ctx.Provider value={statesPage}>
      {children}
    </Job_ctx.Provider>
  )
}
export default JobContextProvider 

