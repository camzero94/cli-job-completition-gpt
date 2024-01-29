export namespace JobType {
  export type Job = {
    jobName: string
    company?: string
    content: string
    link: string
    skills?: string[]
    exp?: string
    date?: string
    location?: string
  }
  export type JobQuery = {
    jobQuery: string,
    skillsQuery: Array<string>,
    page: number | null
  }
}

export default JobType
