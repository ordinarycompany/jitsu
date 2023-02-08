import { useServices } from "../../../hooks/useServices"
import GetStartedPage from "./GetStartedPage"

export default function LoginPage() {
  const services = useServices()
  return (
    <GetStartedPage
      oauthSupport={services.userService.getLoginFeatures().oauth}
      ssoAuthLink={services.userService.getSSOAuthLink()}
      login={true}
      passwordLogin={services.userService.getLoginFeatures().password}
      useCloudHero={services.userService.getLoginFeatures().oauth}
    />
  )
}
