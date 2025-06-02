import type { FC } from 'react'
import { auth } from './firebase'
import { useAuthState } from 'react-firebase-hooks/auth'
import { SignInButton } from './components/atoms/button/SignInButton'
import { SignOutButton } from './components/atoms/button/SignOutButton'
import { UserInfo } from './components/molecules/UserInfo'

export const Home: FC = () => {
  const [user] = useAuthState(auth)

  return (
    <div>
      {user ? (
        <>
          <UserInfo />
          <SignOutButton />
        </>
      ) : (
        <SignInButton />
      )}
    </div>
  )
}
