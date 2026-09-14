
import React, { Fragment } from 'react'
import SpkButton from '../spk-button'
import Link from 'next/link';

interface Collapsecards {
    children?: React.ReactNode;
    footertext?: string;
    Title?: string;
    Timeout?: string;
    Custombodyclass?: string;
}
const SpkCollapsecard: React.FC<Collapsecards> = ({   children, Title, footertext, Timeout, Custombodyclass }) => {
    return (
        <Fragment>
             <div className="box">
                    <div className="box-header justify-between">
                        <div className="box-title">
                                {Title}
                        </div>
                        <Link scroll={false} aria-label="anchor" className="hs-collapse-toggle inline-flex items-center gap-x-2 open" href="#" id="hs-show-hide-collapse" data-hs-collapse="#hs-show-hide-collapse-heading">
                        <svg xmlns="http://www.w3.org/2000/svg"  width="16" height="16" viewBox="0 0 24 24" fill="currentColor"><path d="M11.9999 13.1714L16.9497 8.22168L18.3639 9.63589L11.9999 15.9999L5.63599 9.63589L7.0502 8.22168L11.9999 13.1714Z"></path></svg>
                           </Link> 
                    </div>  
                    <div id="hs-show-hide-collapse-heading" className={`hs-collapse open w-full  overflow-hidden  transition-[height] duration-${Timeout}`} aria-labelledby="hs-show-hide-collapse">
                        <div className={`box-body ${Custombodyclass}`}>
                                {children}
                        </div>
                        <div className="box-footer"> 
                            <SpkButton buttontype="button" variant="primary" customClass="ti-btn">{footertext}</SpkButton> 
                        </div>
                    </div>
            </div>
        </Fragment>
    )
}

export default SpkCollapsecard
