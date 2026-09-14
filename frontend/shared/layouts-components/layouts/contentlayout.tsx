import Footer from '@/shared/layouts-components/footer/footer';
import Header from '@/shared/layouts-components/header/header';
import Sidebar from '@/shared/layouts-components/sidebar/sidebar';
import Switcher from '@/shared/layouts-components/switcher/switcher';
import React, { Fragment } from 'react'
import Backtotop from '@/shared/layouts-components/backtotop/backtotop';
import Loader from '@/shared/layouts-components/loader/loader';
import PrelineScript from '@/pages/PrelineScript';

const Contentlayout = ({ children, }: any) => {

    return (
        <Fragment>
            <Switcher />
            <Loader/>
            <div className='page'>
                <Header />
                <Sidebar />
                <div className='main-content app-content'>
                    <div className='container-fluid'>
                        {children}
                    </div>
                </div>
                <Footer />
            </div>
            <Backtotop />
            <PrelineScript/>
        </Fragment>
    )

}

// const mapStateToProps = (state: any) => ({
//     local_varaiable: state
// });

// export default connect(mapStateToProps, { ThemeChanger })(Contentlayout);
export default Contentlayout;