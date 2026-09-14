
import PrelineScript from '@/pages/PrelineScript';
import LandingSwitcher from '@/shared/layouts-components/switcher/landing-switcher';
import { useRouter } from 'next/router';
import React, { Fragment, useEffect, useRef } from 'react'

const Authenticationlayout = ({ children }: any) => {
 const location = useRouter();
const bodyRef = useRef<any>(null);

  useEffect(() => {
    bodyRef.current = document.body
    // Add 'authentication-background' for authentication-related pages
    if (location.pathname.includes("-basic")) {
      bodyRef.current.classList.add('authentication-background');
    } else {
      bodyRef.current.classList.remove('authentication-background');
    }

    // Add 'coming-soon-main' for the coming soon page
    if (location.pathname.includes('/components/authentication/coming-soon') || location.pathname.includes('/components/authentication/under-maintainance')) {
      bodyRef.current.classList.add('coming-soon-main');
    } else {
      bodyRef.current.classList.remove('coming-soon-main');
    }
    if (location.pathname.includes('/sign-up/sign-up-cover')) {
      bodyRef.current.classList.add('bg-white');
    } else {
      bodyRef.current.classList.remove('bg-white');
    }
    // Clean up classes on route change
    return () => {
      bodyRef.current.classList.remove('authentication-background');
      bodyRef.current.classList.remove('coming-soon-main');
      bodyRef.current.classList.remove('bg-white');
    };
  }, [location.pathname]);

  return (
    <Fragment>
      <LandingSwitcher />
      {children}
      <PrelineScript/>
    </Fragment>
  )

}

export default Authenticationlayout;