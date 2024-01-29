import SearchKeywordComponent from './searchKeywordComponent'
import SkillComponent from './skillComponent'

const SearchComponent: React.FC = () => {


  return (
    <div className='flex  flex-col w-full gap-4'>
      <SearchKeywordComponent />
      <SkillComponent />
    </div>
  )
}

export default SearchComponent
