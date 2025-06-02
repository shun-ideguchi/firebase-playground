import { signInWithPopup } from 'firebase/auth'
import type { FC } from 'react'
import { auth, provider } from '../../../firebase'

export const SignInButton: FC = () => {
  const signInWithGoogle = () => {
    signInWithPopup(auth, provider)
  }

  return (
    <button onClick={signInWithGoogle}>
      <p>グーグルでサインイン</p>
    </button>
  )
}
