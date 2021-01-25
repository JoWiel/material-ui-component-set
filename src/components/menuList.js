(() => ({
  name: 'MenuList',
  type: 'CONTAINER_COMPONENT',
  allowedTypes: ['LIST_ITEM'],
  orientation: 'HORIZONTAL',
  jsx: (() => {
    const { MenuList, List, ListItem, ListItemText } = window.MaterialUI.Core;
    const { env, ModelProvider, useAllQuery, getProperty } = B;
    const isDev = env === 'dev';

    return <MenuList autoFocusItem={!isDev}>{children}</MenuList>;
  })(),
  styles: B => t => {
    const style = new B.Styling(t);
    return {
      root: {
        backgroundColor: ({ options: { backgroundColor } }) =>
          style.getColor(backgroundColor),
      },
      empty: {
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        minHeight: '4rem',
        height: '100%',
        width: '100%',
        fontSize: '0.75rem',
        color: '#262A3A',
        textTransform: 'uppercase',
        boxSizing: 'border-box',
      },
      pristine: {
        borderWidth: '0.0625rem',
        borderColor: '#AFB5C8',
        borderStyle: 'dashed',
        backgroundColor: '#F0F1F5',
        '&::after': {
          content: '"List"',
        },
      },
    };
  },
}))();
