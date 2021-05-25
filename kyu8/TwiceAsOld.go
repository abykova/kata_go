package kata

func TwiceAsOld(dadYearsOld, sonYearsOld int) int { 
  twiceOld := sonYearsOld * 2
  countYears := dadYearsOld - twiceOld
  if twiceOld > dadYearsOld {
    countYears := twiceOld - dadYearsOld
    return countYears
  }
  return countYears
}
