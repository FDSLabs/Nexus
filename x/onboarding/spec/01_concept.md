<!--
order: 1
-->

# Concept

When users onboard to Nexus from Ethereum, they need little amount of Nexus for gas. However, since they don't have any Nexus, they are unable to proceed:

- Create an address on Nexus EVM
- Convert their sdk assets to ERC20 on Nexus EVM
- Swap their Ethereum assets to Nexus for gas usage

To resolve this issue, we need to provide a way for users to swap their Ethereum assets to Nexus automatically.

## Onboarding Process

When users transfer assets to the Nexus network through Gravity Bridge, the IBC transfer automatically triggers swap and conversion to Nexus ERC20 via IBC middleware. These actions are triggered only when transferred through a whitelisted channel.

### Procedure

- User transfers assets to the Nexus network through Gravity Bridge
- Check recipient address's Nexus balance
- If the balance is less than the minimum threshold (currently set to 4), swap the assets to Nexus
- Convert the remaining assets to ERC20

### Middleware ordering

The IBC middleware adds custom logic between the core IBC and the underlying application. Middlewares are implemented as stacks so that applications can define multiple layers of custom behavior.
Function calls from the IBC core to the application travel from the top-level middleware to the bottom middleware, and then to the application.

For Nexus the middleware stack ordering is defined as follows (from top to bottom):

1. IBC Transfer
2. Recovery Middleware
3. Onboarding Middleware

Each module implements their own custom logic in the packet callback `OnRecvPacket`. When a packet arrives from the IBC core, the IBC transfer will be executed first, followed by an attempted recovery, and finally the onboarding will be executed.

### Execution errors

It is possible that the IBC transaction fails in any point of the stack execution and in that case the onboarding will not be triggered by the transaction, as it will rollback to the previous state.
However, the onboarding process is non-atomic, meaning that even if the swap or conversion fails, it does not revert IBC transfer and the asset transferred to the Nexus network will still remain in the Nexus network.
