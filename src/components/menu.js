(() => ({
  name: 'Menu',
  type: 'CONTAINER_COMPONENT',
  allowedTypes: ['LIST_ITEM'],
  orientation: 'VERTICAL',
  jsx: (() => {
    const {
      Button,
      IconButton,
      Menu,
      MenuItem,
      Popper,
      Grow,
      Paper,
      ClickAwayListener,
      MenuList,
    } = window.MaterialUI.Core;
    const { Icons } = window.MaterialUI;

    const {
      variant,
      disabled,
      fullWidth,
      size,
      icon,
      iconPosition,
      linkType,
      linkTo,
      linkToExternal,
      type,
      visible,
      actionId,
      buttonText,
      actionModels,
    } = options;

    const {
      env,
      InteractionScope,
      useText,
      useAction,
      getIdProperty,
      getModel,
      useProperty,
    } = B;
    const isDev = env === 'dev';
    const isIcon = variant === 'icon';
    const buttonContent = useText(buttonText);
    const [isVisible, setIsVisible] = useState(visible);
    const [isOpen, setIsOpen] = useState(false);
    const buttonRef = useRef(null);

    useEffect(() => {
      setIsVisible(visible);
    }, [visible]);

    B.defineFunction('Show', () => setIsVisible(true));
    B.defineFunction('Hide', () => setIsVisible(false));
    B.defineFunction('Show/Hide', () => setIsVisible(s => !s));

    const generalProps = {
      disabled,
      size,
      tabindex: isDev && -1,
    };

    const iconButtonProps = {
      ...generalProps,
      classes: { root: classes.root },
    };

    const buttonProps = {
      ...generalProps,
      fullWidth,
      variant,
      classes: {
        root: classes.root,
        contained: classes.contained,
        outlined: classes.outlined,
      },
      className: !!buttonContent && classes.empty,
      type: isDev ? 'button' : type,
    };
    const compProps = isIcon ? iconButtonProps : buttonProps;
    const ButtonComp = isIcon ? IconButton : Button;

    const handleToggle = () => {
      if (isDev) return;
      setIsOpen(prevOpen => !prevOpen);
    };

    const handleClose = event => {
      if (
        isDev ||
        (buttonRef.current && buttonRef.current.contains(event.target))
      ) {
        return;
      }

      setIsOpen(false);
    };

    const ButtonComponent = (
      <ButtonComp
        ref={buttonRef}
        {...compProps}
        startIcon={
          !isIcon &&
          icon !== 'None' &&
          iconPosition === 'start' &&
          React.createElement(Icons[icon])
        }
        endIcon={
          !isIcon &&
          icon !== 'None' &&
          iconPosition === 'end' &&
          React.createElement(Icons[icon])
        }
        onClick={handleToggle}
      >
        {isIcon &&
          React.createElement(Icons[icon === 'None' ? 'Error' : icon], {
            fontSize: size,
          })}
        {!isIcon && buttonContent}
      </ButtonComp>
    );

    const MenuComp = ({ dev }) => {
      return (
        <>
          {ButtonComponent}
          <Popper
            open={dev || isOpen}
            anchorEl={buttonRef.current}
            role={undefined}
            transition
            disablePortal
          >
            {({ TransitionProps, placement }) => (
              <Grow
                {...TransitionProps}
                style={{
                  transformOrigin:
                    placement === 'bottom' ? 'center top' : 'center bottom',
                }}
              >
                <Paper>
                  <ClickAwayListener onClickAway={handleClose}>
                    {console.log(children)}
                    {children}
                  </ClickAwayListener>
                </Paper>
              </Grow>
            )}
          </Popper>
          {/* <Menu
            anchorEl={buttonRef}
            keepMounted
            open={dev || isOpen}
            onClose={handleClose}
          >
            <MenuItem onClick={handleClose}>Profile</MenuItem>
            <MenuItem onClick={handleClose}>My account</MenuItem>
            <MenuItem onClick={handleClose}>Logout</MenuItem>
          </Menu> */}
        </>
      );
    };

    if (isDev) {
      return (
        <div className={classes.wrapper}>
          <MenuComp dev={isDev} />
        </div>
      );
    }
    return isVisible && <MenuComp />;
  })(),
  styles: B => t => {
    const { env, mediaMinWidth, Styling } = B;
    const style = new Styling(t);
    const getSpacing = (idx, device = 'Mobile') =>
      idx === '0' ? '0rem' : style.getSpacing(idx, device);
    const isDev = env === 'dev';
    return {
      wrapper: {
        display: ({ options: { fullWidth } }) =>
          fullWidth ? 'block' : 'inline-block',
        width: ({ options: { fullWidth } }) => fullWidth && '100%',
        minHeight: '1rem',
        '& > *': {
          pointerEvents: 'none',
        },
      },
      root: {
        color: ({ options: { background, disabled, textColor, variant } }) => [
          !disabled
            ? style.getColor(variant === 'icon' ? background : textColor)
            : 'rgba(0, 0, 0, 0.26)',
          '!important',
        ],
        width: ({ options: { fullWidth, outerSpacing } }) => {
          if (!fullWidth) return 'auto';
          const marginRight = getSpacing(outerSpacing[1]);
          const marginLeft = getSpacing(outerSpacing[3]);
          return `calc(100% - ${marginRight} - ${marginLeft})`;
        },
        marginTop: ({ options: { outerSpacing } }) =>
          getSpacing(outerSpacing[0]),
        marginRight: ({ options: { outerSpacing } }) =>
          getSpacing(outerSpacing[1]),
        marginBottom: ({ options: { outerSpacing } }) =>
          getSpacing(outerSpacing[2]),
        marginLeft: ({ options: { outerSpacing } }) =>
          getSpacing(outerSpacing[3]),
        '&.MuiButton-root, &.MuiIconButton-root': {
          [`@media ${mediaMinWidth(600)}`]: {
            width: ({ options: { fullWidth, outerSpacing } }) => {
              if (!fullWidth) return 'auto';
              const marginRight = getSpacing(outerSpacing[1], 'Portrait');
              const marginLeft = getSpacing(outerSpacing[3], 'Portrait');
              return `calc(100% - ${marginRight} - ${marginLeft})`;
            },
            marginTop: ({ options: { outerSpacing } }) =>
              getSpacing(outerSpacing[0], 'Portrait'),
            marginRight: ({ options: { outerSpacing } }) =>
              getSpacing(outerSpacing[1], 'Portrait'),
            marginBottom: ({ options: { outerSpacing } }) =>
              getSpacing(outerSpacing[2], 'Portrait'),
            marginLeft: ({ options: { outerSpacing } }) =>
              getSpacing(outerSpacing[3], 'Portrait'),
          },
          [`@media ${mediaMinWidth(960)}`]: {
            width: ({ options: { fullWidth, outerSpacing } }) => {
              if (!fullWidth) return 'auto';
              const marginRight = getSpacing(outerSpacing[1], 'Landscape');
              const marginLeft = getSpacing(outerSpacing[3], 'Landscape');
              return `calc(100% - ${marginRight} - ${marginLeft})`;
            },
            marginTop: ({ options: { outerSpacing } }) =>
              getSpacing(outerSpacing[0], 'Landscape'),
            marginRight: ({ options: { outerSpacing } }) =>
              getSpacing(outerSpacing[1], 'Landscape'),
            marginBottom: ({ options: { outerSpacing } }) =>
              getSpacing(outerSpacing[2], 'Landscape'),
            marginLeft: ({ options: { outerSpacing } }) =>
              getSpacing(outerSpacing[3], 'Landscape'),
          },
          [`@media ${mediaMinWidth(1280)}`]: {
            width: ({ options: { fullWidth, outerSpacing } }) => {
              if (!fullWidth) return 'auto';
              const marginRight = getSpacing(outerSpacing[1], 'Desktop');
              const marginLeft = getSpacing(outerSpacing[3], 'Desktop');
              return `calc(100% - ${marginRight} - ${marginLeft})`;
            },
            marginTop: ({ options: { outerSpacing } }) =>
              getSpacing(outerSpacing[0], 'Desktop'),
            marginRight: ({ options: { outerSpacing } }) =>
              getSpacing(outerSpacing[1], 'Desktop'),
            marginBottom: ({ options: { outerSpacing } }) =>
              getSpacing(outerSpacing[2], 'Desktop'),
            marginLeft: ({ options: { outerSpacing } }) =>
              getSpacing(outerSpacing[3], 'Desktop'),
          },
        },
      },
      contained: {
        backgroundColor: ({ options: { background, disabled } }) => [
          !disabled ? style.getColor(background) : 'rgba(0, 0, 0, 0.12)',
          '!important',
        ],
      },
      outlined: {
        borderColor: ({ options: { background, disabled } }) => [
          !disabled ? style.getColor(background) : 'rgba(0, 0, 0, .12)',
          '!important',
        ],
      },
      'MuiPopover-root': {
        position: isDev ? 'relative' : 'fixed',
        zIndex: isDev ? 'unset' : '1300',
      },
      empty: {
        '&::before': {
          content: '"\xA0"',
        },
      },
    };
  },
}))();
