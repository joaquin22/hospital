import  { Fragment, useContext, useEffect } from "react";
import { Helmet } from "react-helmet";
import { connect, useSelector } from "react-redux";
import { createSelector } from "reselect";
import { Initialload } from "../contextapi";
import { ThemeChanger } from "@/shared/redux/action";
import  * as switcherdata from "../../data/switcherdata/switcherdata"
// Create a memoized selector 
const selectRelevantState = createSelector(
  (state) => state, // Input selector
  (state:any) => ({
    dir: state.dir,
    class: state.class,
    dataHeaderStyles: state.dataHeaderStyles,
    dataVerticalStyle: state.dataVerticalStyle,
    dataNavLayout: state.dataNavLayout,
    dataMenuStyles: state.dataMenuStyles,
    toggled: state.toggled,
    dataNavStyle: state.dataNavStyle,
    horStyle: state.horStyle,
    dataPageStyle: state.dataPageStyle,
    dataWidth: state.dataWidth,
    dataMenuPosition: state.dataMenuPosition,
    dataHeaderPosition: state.dataHeaderPosition,
    iconOverlay: state.iconOverlay,
    bgImg: state.bgImg,
    iconText: state.iconText,
    colorPrimaryRgb: state.colorPrimaryRgb,
    PrimaryRgb: state.PrimaryRgb,
    bodyBg: state.bodyBg,
    darkBg: state.darkBg,
    lightRgb: state.lightRgb,
    gray: state.gray,
    inputBorder: state.inputBorder,
  })
);

function RootWrapper({ children , ThemeChanger}:any) {

    const theme = useContext(Initialload);
    useEffect(() => {
      switcherdata.LocalStorageBackup(ThemeChanger, theme.setpageloading);
    }, []);

  // Use the memoized selector
  const local_variable = useSelector(selectRelevantState);

  // Generate dynamic styles as CSS variables
  const customstyles: any = 
  `${local_variable.colorPrimaryRgb != '' ? `--primary-rgb:${local_variable.colorPrimaryRgb}` : ''};
  ${local_variable.PrimaryRgb != '' ? `--primary:${local_variable.PrimaryRgb}` : ''};
  ${local_variable.bodyBg != '' ? `--body-bg:${local_variable.bodyBg}` : ''};
  ${local_variable.darkBg != '' ? `--dark-bg:${local_variable.darkBg}` : ''};
  ${local_variable.lightRgb != '' ? `--light:${local_variable.darkBg}` : ''};
  ${local_variable.gray != '' ? `--gray-3:${local_variable.gray}` : ''};
  ${local_variable.inputBorder != '' ? `--input-border:${local_variable.inputBorder}` : ''};`;

  

  return (
    <Fragment>
        <Helmet>
        <html
            dir={local_variable.dir}
            className={local_variable.class}
            data-header-styles={local_variable.dataHeaderStyles}
            data-vertical-style={local_variable.dataVerticalStyle}
            data-nav-layout={local_variable.dataNavLayout}
            data-menu-styles={local_variable.dataMenuStyles}
            data-toggled={local_variable.toggled}
            data-nav-style={local_variable.dataNavStyle}
            data-page-style={local_variable.dataPageStyle}
            data-width={local_variable.dataWidth}
            data-menu-position={local_variable.dataMenuPosition}
            data-header-position={local_variable.dataHeaderPosition}
            data-icon-overlay={local_variable.iconOverlay}
            bg-img={local_variable.bgImg}
            icon-text={local_variable.iconText}
            style={customstyles} // Apply styles directly
              >
              </html>
        </Helmet>
        {children}
    </Fragment>
  );
}

const mapStateToProps = (state: any) => ({
  local_varaiable: state
});

export default connect(mapStateToProps, { ThemeChanger })(RootWrapper);
// export default RootWrapper;
